package service

import (
	"gozealous/errorr"
	"gozealous/service/native"
	"gozealous/service/railway"
)

type Native interface {
	SystemHealth() *native.SystemState
	DatabaseHealth() *native.DatabaseState
}

type Railway interface {
	Stations() (interface{}, errorr.Entity)
	Lines() ([]railway.Line, errorr.Entity)
	Journey(originStationName string, destinationStationName string) ([][]railway.PathPoint, errorr.Entity)
}
