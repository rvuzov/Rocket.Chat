package rocketchat

type apiResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorType string `json:"errorType"`
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
