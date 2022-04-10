package repository

import (
	"gozealous/anomaly"
	"gozealous/repository/railway"
)

type Database interface {
	Ping() *anomaly.ServiceError
}

type Railway interface {
	Stations() ([]railway.Station, *anomaly.ServiceError)
}
