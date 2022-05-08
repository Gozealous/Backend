package common

import (
	"fmt"
	"gozealous/errorr"
)

func ParseError(code string, message string, err error) errorr.Entity {
	errMsg := fmt.Sprintf("%s %s", message, err.Error())

	return errorr.New(code, errMsg, nil)
}
