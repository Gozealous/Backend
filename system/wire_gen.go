// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package system

import (
	"database/sql"
	"github.com/google/wire"
	"gozealous/log"
	"gozealous/nexus"
	"gozealous/repository"
	"gozealous/repository/configuration"
	"gozealous/repository/database"
	"gozealous/repository/railway"
	"gozealous/service"
	"gozealous/service/native"
	railway2 "gozealous/service/railway"
)

// Injectors from wire.go:

func InitialiseNexus(logger log.Logger, db *sql.DB) *nexus.Store {
	repository := database.NewRepository(db)
	service := native.NewService(logger, repository)
	configurationRepository := configuration.NewRepository(db)
	railwayRepository := railway.NewRepository(db)
	railwayService := railway2.NewService(logger, configurationRepository, railwayRepository)
	store := nexus.NewStore(service, railwayService)
	return store
}

// wire.go:

var databaseRepositorySet = wire.NewSet(database.NewRepository, wire.Bind(new(repository.Database), new(*database.Repository)))

var configurationRepositorySet = wire.NewSet(configuration.NewRepository, wire.Bind(new(repository.Configuration), new(*configuration.Repository)))

var railwayRepositorySet = wire.NewSet(railway.NewRepository, wire.Bind(new(repository.Railway), new(*railway.Repository)))

var nativeServiceSet = wire.NewSet(native.NewService, wire.Bind(new(service.Native), new(*native.Service)))

var railwayServiceSet = wire.NewSet(railway2.NewService, wire.Bind(new(service.Railway), new(*railway2.Service)))
