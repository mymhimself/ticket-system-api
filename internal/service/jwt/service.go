package jwt

import (
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	service2 "github.com/mymhimself/ticket-system-api/internal/service"
)

type jwtService struct {
	cfg    *config.Authentication
	logger *logger.Logger
}

func New(cfg config.Authentication, logger logger.Logger) service2.Authentication {
	return jwtService{
		cfg:    &cfg,
		logger: &logger,
	}
}
