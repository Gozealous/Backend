package railway

import (
	"gozealous/log"
	"gozealous/repository"
)

type Service struct {
	logger      log.Logger
	railwayRepo repository.Railway
}

func NewService(logger log.Logger, railwayRepository repository.Railway) *Service {
	return &Service{
		logger:      logger,
		railwayRepo: railwayRepository,
	}
}
