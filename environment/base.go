package environment

import (
	"fmt"
	"gozealous/code"
	"gozealous/errorr"
	"os"
)

func getValue(key string) (string, errorr.Entity) {
	value, found := os.LookupEnv(key)

	if !found {
		err := errorr.New(
			code.EnvironmentVariableNotFound,
			fmt.Sprintf("%s not found", key),
			nil,
		)
		return "", err
	}

	return value, nil
}
