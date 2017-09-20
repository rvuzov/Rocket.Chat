package rocketchat

import (
	"fmt"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("rocketchat")

func makeError(err *error, functionName string, args ...interface{}) {
	if err == nil || *err == nil {
		return
	}
	functionArgs := ""
	if len(args)%2 == 1 {
		args = append(args, "")
	}
	for i := 0; i < len(args); i += 2 {
		functionArgs += "%v=%v; "
	}
	format := fmt.Sprintf("%s(%s): %s", functionName, functionArgs, (*err).Error())
	*err = fmt.Errorf(format, args...)
}

