package main

import (
	"../../chat"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	//CreateUserToken()
	//SetUserAvatar()
	UpdateUser()
}

func CreateUser() {
	userID, err := rocketchat.Users.Create("bla@gmail.com", "1111", "Oleh Lamzin", "lamzin")
	if err != nil {
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success: %s", userID)
	}
}

func CreateUserToken() {
	userID, err := rocketchat.Users.CreateToken("XqCkNXQYwRXm5QfhK")
	if err != nil {
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success: %s", userID)
	}
}

func SetUserAvatar() {
	err := rocketchat.Users.SetAvatar("XqCkNXQYwRXm5QfhK", "https://www.smashingmagazine.com/wp-content/uploads/2015/06/10-dithering-opt.jpg")
	if err != nil {
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success")
	}
}

func UpdateUser() {
	err := rocketchat.Users.Update("XqCkNXQYwRXm5QfhK", "Vasya")
	if err != nil {
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success")
	}
}
