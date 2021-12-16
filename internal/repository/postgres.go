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
	SaveTicket(ticket *model.Ticket) error
	SaveTicketReply(reply *model.Ticket) error
	//load tickets for specific user and specific status type
	LoadTickets(userID uint, status enum.TicketStatus) ([]model.Ticket, error)
	LoadAllTickets(status enum.TicketStatus) ([]model.Ticket, error)
}
