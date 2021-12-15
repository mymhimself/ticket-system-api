package service

import "github.com/mymhimself/ticket-system-api/internal/entity/model"

type Authentication interface {
	GenerateToken(user *model.User) (string, error)
	GenerateRefreshToken(user *model.User) (string, error)
}
