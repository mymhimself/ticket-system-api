package service

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Authentication interface {
	GenerateTokenPair(user *model.User) (string, string, error)
}
