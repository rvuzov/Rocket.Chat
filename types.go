package rocketchat

import "fmt"

type apiResponseInterface interface {
	GetError() error
}

type apiResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorType string `json:"errorType"`
}

func (r apiResponse) GetError() error {
	if r.Success {
		return nil
	}
	return fmt.Errorf("%s", r.Error)
}

type apiUser struct {
	Id string `json:"_id"`
}

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

type (
	apiSetUserAvatarRequest struct {
		UserId    string `json:"userId"`
		AvatarUrl string `json:"avatarUrl"`
	}

	apiSetUserAvatarResponse struct {
		apiResponse
	}
)

type (
	apiUpdateUserRequest struct {
		UserId string                   `json:"userId"`
		Data   apiUpdateUserRequestData `json:"data"`
	}

	apiUpdateUserRequestData struct {
		Name string `json:"name"`
	}

	apiUpdateUserResponse struct {
		apiResponse
		User apiUser `json:"user"`
	}
)
