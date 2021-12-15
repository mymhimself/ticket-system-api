package account

import (
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/repository"
	"github.com/mymhimself/ticket-system-api/internal/service"
)

type accountService struct {
	cfg            config.Account
	postgres       repository.Postgres
	Authentication service.Authentication
	logger         logger.Logger
}

func New(cfg config.Account, auth service.Authentication, postgres repository.Postgres, logger logger.Logger) service.Account {
	return &accountService{
		cfg:            cfg,
		postgres:       postgres,
		logger:         logger,
		Authentication: auth,
	}
}
