package validation

import (
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/service"
)

type validationImpl struct {
	//cfg        config.Validation
	logger logger.Logger
}

func New(logger logger.Logger) service.Validation {
	return &validationImpl{
		logger: logger,
	}
}
