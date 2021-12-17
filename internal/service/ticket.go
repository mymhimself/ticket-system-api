package service

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Ticket interface {
	CreateNewTicket(ticket *model.TicketThread, text string) error
	LoadTicketByAdmin(status enum.TicketThreadStatus) ([]model.Ticket, error)
	LoadTicketByUser(status enum.TicketThreadStatus, userID uint) ([]model.Ticket, error)
	GetTicketByID(ticketID uint) (*model.Ticket, error)
	ReplyMessage(text string, repliedTicketID uint, senderID uint) error
	ReplyMessageAsAdmin(text string, repliedTicketID uint, senderID uint) error
	UpdateTicket(ticket *model.Ticket) error
	GetFilteredTicketThreads(creatorID uint, seen *bool, status *int64, replied *bool, priority *int64, department *int64) ([]model.TicketInfo, error)
}
