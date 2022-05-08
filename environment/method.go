package environment

import (
	"gozealous/errorr"
)

func ServerPort() (string, errorr.Entity) {
	return getValue("PORT")
}

func DatabaseUri() (string, errorr.Entity) {
	return getValue("DATABASE_URL")
}

func LogLevel() (string, errorr.Entity) {
	return getValue("LOG_LEVEL")
}
