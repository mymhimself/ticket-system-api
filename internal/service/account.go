package service

import (
	"context"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Account interface {
	Login(ctx context.Context, username, password string) (string, error)
	Register(ctx context.Context, user *model.User) (string, error)
	RegisterConfirm(ctx context.Context, key, phone, otp string) (string, error)
}
