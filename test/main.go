package main

import (
	"../../chat"
	"github.com/op/go-logging"
	//"github.com/go-errors/errors"
)

var log = logging.MustGetLogger("main")

func main() {
	userID, err := rocketchat.Users.CreateUser("bla@gmail.com", "1111", "Oleh Lamzin", "lamzin")
	if err != nil {
		//log.Error("Error:", err.(*errors.Error).ErrorStack())
		log.Error("Error:", err.Error())
	} else {
		log.Infof("Success: %s", userID)
	}
}
