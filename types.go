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

func (response apiResponse) GetError() error {
	if response.Success {
		return nil
	}
	return fmt.Errorf("%s", response.Error)
}

type apiUser struct {
	Id string `json:"_id"`
}
