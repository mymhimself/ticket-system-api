package service

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Ticket interface {
	CreateNewTicket(ticket *model.TicketThread, text string) error
	LoadTicketByAdmin(status enum.TicketThreadStatus) ([]model.Ticket, error)
	LoadTicketByUser(status enum.TicketThreadStatus, userID uint) ([]model.Ticket, error)
	GetTicketByID(ticketID uint) (*model.Ticket, bool, error)
	ReplyMessage(text string, repliedTicketID uint, senderID uint) error
	ReplyMessageByAdmin(text string, repliedTicketID uint, senderID uint) error
	UpdateTicketMessage(ticket *model.TicketMessage) error
	GetFilteredTicketThreads(creatorID *uint, seen *bool, status *enum.TicketThreadStatus, replied *bool, priority *enum.Priority, department *int64) ([]model.TicketInfo, error)
	ChangeTicketStatus(ticketID uint, status enum.TicketThreadStatus) error
}
