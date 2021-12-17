package postgres

import (
	"errors"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
	"gorm.io/gorm"
)

func (p *postgres) CreateMessage(message *model.TicketMessage) error {
	result := p.db.Create(message)
	if result.Error != nil {
		p.logger.Error(result.Error.Error())
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		return nil
	}
}

func (p *postgres) UpdateMessage(ticketMessage *model.TicketMessage) error {
	result := p.db.Save(ticketMessage)
	if result.Error != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	}
	return nil
}

func (p *postgres) CreateMessageWithReply(message *model.TicketMessage, repliedMessage *model.TicketMessage) error {
	tx := p.db.Begin()
	result := tx.Create(message)
	if result.Error != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		result = tx.Save(repliedMessage)
		if result.Error != nil {
			tx.Rollback()
			return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
		} else {
			tx.Commit()
			return nil
		}
	}
}

func (p *postgres) CreateNewTicketThread(firstMessage *model.TicketMessage) error {
	result := p.db.Create(firstMessage)
	if result.Error != nil {
		p.logger.Error(result.Error.Error())
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		return nil
	}
}

func (p *postgres) GetUserTickets(userID uint, status enum.TicketThreadStatus) ([]model.Ticket, error) {
	var tickets []model.Ticket
	result := p.db.Where("sender_id = ? and status = ?", userID, status).Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return tickets, nil
	}
}

func (p *postgres) GetAllTickets(status enum.TicketThreadStatus) ([]model.Ticket, error) {
	var tickets []model.Ticket
	result := p.db.Where(model.TicketThread{Status: status}).Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return tickets, nil
	}
}

func (p *postgres) UpdateTicket(ticket *model.Ticket) error {
	result := p.db.Create(ticket)
	if result.Error != nil {
		return myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
	} else {
		return nil
	}
}

func (p *postgres) GetTicketMessageByID(id uint) (*model.TicketMessage, bool, error) {
	var message model.TicketMessage
	result := p.db.First(&message, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, nil
		} else {
			return nil, false, myerror.New(myerror.InternalError, enum.RepoLayer, result.Error.Error())
		}
	} else {
		return &message, true, nil
	}
}

func (p *postgres) GetAdminTicketThreadList(userID uint, isOpen bool) ([]model.TicketThread, error) {
	return nil, nil
}

func (p *postgres) GetUserTicketThreadList(userID uint) ([]model.TicketThread, error) {
	return nil, nil
}
