package environment

import (
	"fmt"
	"gozealous/anomaly"
	"gozealous/code"
	"os"
)

func getValue(key string) (string, error) {
	value, found := os.LookupEnv(key)

	if !found {
		err := &anomaly.ServiceError{
			Code:    code.EnvironmentVariableNotFound,
			Message: fmt.Sprintf("%s not found", key),
		}
		return "", err
	}

	return value, nil
}
