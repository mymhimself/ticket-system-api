package repository

import (
	"context"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Postgres interface {
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	DeleteUser(ctx context.Context, id int) error
	//
	SendTicket(ctx context.Context, ticket *model.Ticket, user *model.User) error
}
