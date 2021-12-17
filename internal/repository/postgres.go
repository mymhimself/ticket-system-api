package repository

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

type Postgres interface {
	AutoMigration() error
	//user relevant methods
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	DeleteUser(id uint) error
	CheckUserExistence(username string) (bool, error)

	//ticket relevant methods
	//load tickets for specific user and specific status type
	CreateNewTicketThread(firstMessage *model.TicketMessage) error
	CreateMessage(ticketMessage *model.TicketMessage) error
	CreateMessageWithReply(message *model.TicketMessage, repliedMessage *model.TicketMessage) error
	UpdateMessage(ticketMessage *model.TicketMessage) error

	UpdateTicket(reply *model.Ticket) error
	GetUserTickets(userID uint, status enum.TicketThreadStatus) ([]model.Ticket, error)
	GetAllTickets(status enum.TicketThreadStatus) ([]model.Ticket, error)
	GetTicketMessageByID(id uint) (*model.TicketMessage, bool, error)
	GetOpenTicketThreadsList() ([]model.TicketThread, error)
	GetUserTicketThreadListByFilter(userID uint, filterParams map[string]interface{}) ([]model.TicketInfo, error)
}
