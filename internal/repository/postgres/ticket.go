package postgres

import (
	"context"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func (p *postgres) SendTicket(ctx context.Context, ticket *model.Ticket, user *model.User) error {
	return nil
}
