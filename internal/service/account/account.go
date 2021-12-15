package account

import (
	"context"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func (a *accountService) Login(ctx context.Context, username, password string) (string, error) {
	// start span with context

	user, err := a.postgres.GetUserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	_ = user
	// check password

	// create token
	token := ""

	return token, nil
}

func (a *accountService) Register(ctx context.Context, user *model.User) (string, error) {
	return "", nil
}

func (a *accountService) RegisterConfirm(ctx context.Context, key, phone, otp string) (string, error) {
	return "", nil
}
