package ticket

import (
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/repository"
	"github.com/mymhimself/ticket-system-api/internal/service"
)

type ticketImpl struct {
	//cfg        config.Validation
	logger   logger.Logger
	postgres repository.Postgres
}

func New(logger logger.Logger, postgres repository.Postgres) service.Ticket {
	return &ticketImpl{
		logger:   logger,
		postgres: postgres,
	}
}
