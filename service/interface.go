package service

import (
	"gozealous/anomaly"
	"gozealous/service/native"
)

type Native interface {
	SystemHealth() *native.SystemState
	DatabaseHealth() *native.DatabaseState
}

type Railway interface {
	Stations() (interface{}, *anomaly.ServiceError)
}
