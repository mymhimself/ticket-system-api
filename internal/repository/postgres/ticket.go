package postgres

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func (p *postgres) SaveTicket(ticket *model.Ticket) error {
	result := p.db.Create(ticket)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (p *postgres) LoadTickets(userID uint, status enum.TicketStatus) ([]model.Ticket, error) {
	var tickets []model.Ticket
	result := p.db.Where("sender_id = ? and status = ?", userID, status).Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return tickets, nil
	}
}

func (p *postgres) LoadAllTickets(status enum.TicketStatus) ([]model.Ticket, error) {
	var tickets []model.Ticket
	result := p.db.Where(model.Ticket{Status: status}).Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return tickets, nil
	}
}

func (p *postgres) SaveTicketReply(reply *model.Ticket) error {
	result := p.db.Create(reply)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
