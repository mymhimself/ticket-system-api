package ticket

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func (s ticketImpl) SendTicketByUser(ticket *model.Ticket) error {
	if err := s.postgres.SaveTicket(ticket); err != nil {
		return err
	}
	return nil
}

func (s ticketImpl) LoadTicketByAdmin(status enum.TicketStatus) ([]model.Ticket, error) {
	return s.postgres.LoadAllTickets(status)
}
func (s ticketImpl) LoadTicketByUser(status enum.TicketStatus, userID uint) ([]model.Ticket, error) {
	return s.postgres.LoadTickets(userID, status)
}

func (s ticketImpl) ReplyTicket(reply *model.Ticket) error {
	if err := s.postgres.SaveTicket(reply); err != nil {
		return err
	}
	return nil
}
