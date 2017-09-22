package rocketchat

var Users users

type users struct{}

func (users) Create(email, password, name, username string) (userId string, err error) {
	defer makeError(&err, "Create", "email", email, "password", password, "name", name, "username", username)

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

	userId = response.User.Id
	return
}

func (users) CreateToken(userId string) (token string, err error) {
	defer makeError(&err, "CreateToken", "userId", userId)

	request := apiCreateUserTokenRequest{UserId: userId}
	var response apiCreateUserTokenResponse
	if err = apiPost(kapiUsersCreateToken, request, &response); err != nil {
		return
	}

	token = response.Data.AuthToken
	return
}

func (users) SetAvatar(userId string, avatarUrl string) (err error) {
	defer makeError(&err, "SetAvatar", "userId", userId, "avatarUrl", avatarUrl)

	request := apiSetUserAvatarRequest{UserId: userId, AvatarUrl: avatarUrl}
	var response apiSetUserAvatarResponse
	err = apiPost(kapiUsersSetAvatar, request, &response)
	return
}

func (users) Update(userId string, name string) (err error) {
	defer makeError(&err, "Update", "userId", userId, "name", name)

	request := apiUpdateUserRequest{
		UserId: userId,
		Data:   apiUpdateUserRequestData{Name: name},
	}
	var response apiUpdateUserResponse
	err = apiPost(kapiUsersUpdate, request, &response)
	return
}
