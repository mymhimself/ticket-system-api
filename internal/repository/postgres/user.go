package postgres

import (
	"context"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func (p *postgres) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (p *postgres) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (p *postgres) GetUserByID(ctx context.Context, id int) (*model.User, error) {
	return nil, nil
}

func (p *postgres) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return nil, nil
}

func (p *postgres) DeleteUser(ctx context.Context, id int) error {
	return nil
}
