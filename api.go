package rocketchat

import (
	"bytes"
	"encoding/json"
	"github.com/rvuzov/goapp-config"
	"io/ioutil"
	"net/http"
)

func apiPost(method apiMethod, request interface{}, response apiResponseInterface) (err error) {
	defer makeError(&err, "apiPost", "method", method, "request", request, "response", response)

	requestJson, err := json.Marshal(request)
	if err != nil {
		return
	}

	httpRequest, err := http.NewRequest(
		"POST",
		config.Get("ROCKET_CHAT_URL")+string(method),
		bytes.NewReader(requestJson),
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
	//log.Debug(string(responseBody))

	if err = json.Unmarshal(responseBody, response); err != nil {
		return
	}

	err = response.GetError()
	return
}
