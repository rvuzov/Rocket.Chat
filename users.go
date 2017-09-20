package rocketchat

import "fmt"

var Users users

type users struct{}

func (users) CreateUser(email, password, name, username string) (userId string, err error) {
	defer makeError(&err, "CreateUser", "email", email, "password", password, "name", name, "username", username)

	request := apiCreateUserRequest{
		Email:               email,
		Password:            password,
		Name:                name,
		UserName:            username,
		JoinDefaultChannels: false,
	}

	var response apiCreateUserResponse
	if err = apiPost(kapiUsersCreate, request, &response); err != nil {
		return
	}

	if !response.Success {
		err = fmt.Errorf("%s", response.Error)
		return
	}

	userId = response.User.Id
	return
}
