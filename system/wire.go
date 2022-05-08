//go:build wireinject
// +build wireinject

package system

import (
	"database/sql"
	"gozealous/log"
	"gozealous/nexus"
	"gozealous/repository"
	configRepo "gozealous/repository/configuration"
	databaseRepo "gozealous/repository/database"
	railwayRepo "gozealous/repository/railway"
	"gozealous/service"
	"gozealous/service/native"
	"gozealous/service/railway"

	"github.com/google/wire"
)

var databaseRepositorySet = wire.NewSet(
	databaseRepo.NewRepository,
	wire.Bind(new(repository.Database), new(*databaseRepo.Repository)),
)

var configurationRepositorySet = wire.NewSet(
	configRepo.NewRepository,
	wire.Bind(new(repository.Configuration), new(*configRepo.Repository)),
)

var railwayRepositorySet = wire.NewSet(
	railwayRepo.NewRepository,
	wire.Bind(new(repository.Railway), new(*railwayRepo.Repository)),
)

var nativeServiceSet = wire.NewSet(
	native.NewService,
	wire.Bind(new(service.Native), new(*native.Service)),
)

var railwayServiceSet = wire.NewSet(
	railway.NewService,
	wire.Bind(new(service.Railway), new(*railway.Service)),
)

func InitialiseNexus(logger log.Logger, db *sql.DB) *nexus.Store {
	panic(
		wire.Build(
			databaseRepositorySet,
			configurationRepositorySet,
			railwayRepositorySet,

			nativeServiceSet,
			railwayServiceSet,

			nexus.NewStore,
		),
	)
}
