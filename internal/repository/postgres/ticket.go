package postgres

import (
	"errors"
	"fmt"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
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
	result := p.db.Preload(clause.Associations).First(&message, id)
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

func (p *postgres) GetOpenTicketThreadsList() ([]model.TicketThread, error) {
	var threads []model.TicketThread
	res := p.db.Where(model.TicketThread{Status: enum.Active}).Find(&threads)
	if res.Error != nil {
		return nil, myerror.New(myerror.InternalError, enum.RepoLayer, res.Error.Error())
	} else if res.RowsAffected == 0 {
		return nil, nil
	} else {
		return threads, nil
	}
}

func (p *postgres) GetUserTicketThreadListByFilter(userID uint, filterParams map[string]interface{}) ([]model.TicketInfo, error) {
	var list []model.TicketInfo
	var conditions []string

	query := fmt.Sprintf(QryTicketInfoFilter, userID)
	for k, v := range filterParams {
		switch k {
		case "status":
			conditions = append(conditions, fmt.Sprintf("status = %d", v.(int64)))
		case "department":
			conditions = append(conditions, fmt.Sprintf("department = %d", v.(int64)))
		case "seen":
			conditions = append(conditions, fmt.Sprintf("bool(is_seen) = %v", v.(bool)))
		case "replied":
			conditions = append(conditions, fmt.Sprintf("bool(is_replied) = %v", v.(bool)))
		case "priority":
			conditions = append(conditions, fmt.Sprintf("priority = %d", v.(int64)))
		}
	}
	if len(conditions) != 0 {
		query = query + " where " + strings.Join(conditions, " and ")
	}
	fmt.Println(query)

	res := p.db.Raw(query).Find(&list)
	if res.Error != nil {
		return nil, res.Error
	} else {
		return list, nil
	}
}

func (p *postgres) GetAdminTicketThreadsList(adminID uint) {

}
