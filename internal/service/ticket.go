package service

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Ticket interface {
	SendTicketByUser(ticket *model.Ticket) error
	LoadTicketByAdmin(status enum.TicketStatus) ([]model.Ticket, error)
	LoadTicketByUser(status enum.TicketStatus, userID uint) ([]model.Ticket, error)
	ReplyTicket(reply *model.Ticket) error
}
