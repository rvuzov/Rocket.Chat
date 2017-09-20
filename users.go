package rocketchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AlexeySpiridonov/goapp-config"
	"github.com/go-errors/errors"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
)

type rocketAPIMethod string

const (
	usersCreate rocketAPIMethod = "/api/v1/users.create"
)

type users struct{}

var Users users

var log = logging.MustGetLogger("rocketchat")

type rocketAPIResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorType string `json:"errorType"`
}

type rocketAPIUser struct {
	Id string `json:"_id"`
}

type rocketAPICreateUserResponse struct {
	rocketAPIResponse
	User rocketAPIUser `json:"user"`
}

func (users) CreateUser(email, password, name, username string) (userId string, err error) {
	defer makeError(&err, "CreateUser", "email", email, "password", password, "name", name, "username", username, 12)

	jsonBody := fmt.Sprintf(
		`{"email": "%s", "password": "%s", "name": "%s", "username": "%s", "joinDefaultChannels": false}`,
		email, password, name, username,
	)
	httpRequest, err := http.NewRequest(
		"POST",
		config.Get("ROCKET_CHAT_URL")+"/api/v1/users.create",
		bytes.NewReader([]byte(jsonBody)),
	)
	httpRequest.Header.Add("X-User-Id", config.Get("ROCKET_CHAT_ADMIN_USER_ID"))
	httpRequest.Header.Add("X-Auth-Token", config.Get("ROCKET_CHAT_ADMIN_AUTH_TOKEN"))
	httpRequest.Header.Add("Content-type", config.Get("application/json"))

	if err != nil {
		return
	}

	client := &http.Client{}
	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return
	}
	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}

	log.Info(string(responseBody))

	var response rocketAPICreateUserResponse
	if err = json.Unmarshal(responseBody, &response); err != nil {
		return
	}

	if !response.Success {
		err = errors.Errorf("%s", response.Error)
		return
	}

	userId = response.User.Id
	return
}

func apiPost(method rocketAPIMethod, json string, result interface{}) (err error) {

}

func makeError(err *error, functionName string, args ...interface{}) {
	if err == nil || *err == nil {
		return
	}
	functionArgs := ""
	if len(args)%2 == 1 {
		args = append(args, "")
	}
	for i := 0; i < len(args); i += 2 {
		functionArgs += "%v=%v; "
	}
	format := fmt.Sprintf("rocketchat: %s(%s): %s", functionName, functionArgs, (*err).Error())
	*err = errors.Errorf(format, args...)
}
