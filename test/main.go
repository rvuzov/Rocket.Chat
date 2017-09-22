package main

import (
	"../../chat"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func main() {
	//CreateUserToken()
	//SetUserAvatar()
	//UpdateUser()
	SpeedTestPostMessage()
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

func SpeedTestPostMessage() {
	roomId := []string{
		"YqLot5EBGWGHkKHvdojCuFJrmHYn6dzNDX",
		"general",
	}

	wait := make(chan struct{}, 1)

	goroutines := 100
	messagesPerGoroutine := 100

	for i := 0; i < goroutines; i++ {
		go func(index int) {
			for j := 0; j < messagesPerGoroutine; j++ {
				PostMessage(roomId[index % len(roomId)])
			}
			wait <- struct{}{}
		}(i)
	}

	for i := 0; i < goroutines; i++ {
		<-wait
	}

}

func PostMessage(roomId string) {
	err := rocketchat.Channel.PostMessage(roomId, "привет!")
	if err != nil {
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success")
	}
}
