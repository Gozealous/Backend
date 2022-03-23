package database

import (
	"fmt"
	"gozealous/anomaly"
)

func parseError(code string, message string, err error) *anomaly.ServiceError {
	errMsg := fmt.Sprintf("%s %s", message, err.Error())

	return &anomaly.ServiceError{
		Code:    code,
		Message: errMsg,
	}
}
