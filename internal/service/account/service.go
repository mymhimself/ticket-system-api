package account

import (
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/repository"
	"github.com/mymhimself/ticket-system-api/internal/service"
)

type accountService struct {
	postgres       repository.Postgres
	Authentication service.Authentication
	logger         logger.Logger
}

func New(auth service.Authentication, postgres repository.Postgres, logger logger.Logger) service.Account {
	return &accountService{
		postgres:       postgres,
		logger:         logger,
		Authentication: auth,
	}
}
