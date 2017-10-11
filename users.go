package rocketchat

var Users users

type users struct{}

type (
	apiCreateUserRequest struct {
		Email               string `json:"email"`
		Password            string `json:"password"`
		Name                string `json:"name"`
		UserName            string `json:"username"`
		JoinDefaultChannels bool   `json:"joinDefaultChannels"`
	}

	apiCreateUserResponse struct {
		apiResponse
		User apiUser `json:"user"`
	}
)

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

type (
	apiCreateUserTokenRequest struct {
		UserId string `json:"userId"`
	}

	apiCreateUserTokenResponseData struct {
		UserId    string `json:"userId"`
		AuthToken string `json:"authToken"`
	}

	apiCreateUserTokenResponse struct {
		apiResponse
		Data apiCreateUserTokenResponseData `json:"data"`
	}
)

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

type (
	apiSetUserAvatarRequest struct {
		UserId    string `json:"userId"`
		AvatarUrl string `json:"avatarUrl"`
	}

	apiSetUserAvatarResponse struct {
		apiResponse
	}
)

func (users) SetAvatar(userId string, avatarUrl string) (err error) {
	defer makeError(&err, "SetAvatar", "userId", userId, "avatarUrl", avatarUrl)

	request := apiSetUserAvatarRequest{UserId: userId, AvatarUrl: avatarUrl}
	var response apiSetUserAvatarResponse
	err = apiPost(kapiUsersSetAvatar, request, &response)
	return
}

type (
	apiUpdateUserRequest struct {
		UserId string                   `json:"userId"`
		Data   apiUpdateUserRequestData `json:"data"`
	}

	apiLoginRequest struct {
		User       string `json:"user"`
		Password   string `json:"password"`
	}

	apiUpdateUserRequestData struct {
		Name string `json:"name"`
	}

	apiUpdateUserResponse struct {
		apiResponse
		User apiUser `json:"user"`
	}

	apiLoginResponse struct {
		apiResponse
		Data authData `json:"data"`
	}
)

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
func (users) Login(userEmail string, userPassword string) (authToken string,userId string,err error) {
	defer makeError(&err, "Login", "userEmail", userEmail, "userPassword", userPassword)

	request := apiLoginRequest{
		User: userEmail,
		Password:userPassword,
	}
	var response apiLoginResponse
	if err = apiPost(kapiUsersLogin, request, &response); err != nil {
		return
	}

	authToken = response.Data.AuthToken
	userId = response.Data.UserId
	return
}