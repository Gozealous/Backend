package database

import (
	"database/sql"
	"gozealous/code"
	"gozealous/environment"
	"gozealous/errorr"
	"gozealous/repository/common"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, errorr.Entity) {
	uri, uriErr := environment.DatabaseUri()
	if uriErr != nil {
		return nil, uriErr.Trace()
	}

	db, sqlOpenErr := sql.Open("postgres", uri)
	if sqlOpenErr != nil {
		err := common.ParseError(code.DatabaseConnectionFailure, "Unable to open connection to database.", sqlOpenErr)
		return nil, err.Trace()
	}

	return db, nil
}
