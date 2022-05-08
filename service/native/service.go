package native

import (
	"gozealous/log"
	"gozealous/repository"
)

func NewService(logger log.Logger, databaseRepository repository.Database) *Service {
	return &Service{
		logger:       logger,
		databaseRepo: databaseRepository,
	}
}
