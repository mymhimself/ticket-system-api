package ticket

import (
	"fmt"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/pkg/myerror"
)

func (ts ticketImpl) CreateNewTicket(thread *model.TicketThread, text string) error {
	var message model.TicketMessage
	message.Text = text
	message.Thread = *thread
	message.ThreadID = thread.ID
	message.Seen = false

	if err := ts.postgres.CreateNewTicketThread(&message); err != nil {
		return err
	}
	return nil
}

func (ts ticketImpl) LoadTicketByAdmin(status enum.TicketThreadStatus) ([]model.Ticket, error) {
	return ts.postgres.GetAllTickets(status)
}

func (ts ticketImpl) LoadTicketByUser(status enum.TicketThreadStatus, userID uint) ([]model.Ticket, error) {
	return ts.postgres.GetUserTickets(userID, status)
}

func (ts ticketImpl) ReplyMessage(text string, repliedMessageID uint, senderID uint) error {
	repliedMessage, exist, err := ts.postgres.GetTicketMessageByID(repliedMessageID)
	if err != nil {
		return err
	} else if !exist {
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
	}
	if repliedMessage.Thread.Status == enum.Closed {
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
	}

	var newMessage model.TicketMessage
	newMessage.ReplyTo = repliedMessage.ID
	newMessage.ThreadID = repliedMessage.ThreadID
	newMessage.Text = text

	repliedMessage.Seen = true
	repliedMessage.Replied = true
	if err := ts.postgres.CreateMessageWithReply(&newMessage, repliedMessage); err != nil {
		return err
	}

	return nil
}

func (ts ticketImpl) ReplyMessageAsAdmin(text string, repliedMessageID uint, senderID uint) error {
	repliedMessage, exist, err := ts.postgres.GetTicketMessageByID(repliedMessageID)
	if err != nil {
		return err
	} else if !exist {
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
	}

	fmt.Println("**", repliedMessage.Thread)

	var newMessage model.TicketMessage
	newMessage.ReplyTo = repliedMessage.ID
	newMessage.ThreadID = repliedMessage.ThreadID
	newMessage.Text = text

	repliedMessage.Seen = true
	repliedMessage.Replied = true
	if err := ts.postgres.CreateMessageWithReply(&newMessage, repliedMessage); err != nil {
		return err
	}

	return nil
}

func (ts ticketImpl) UpdateTicket(ticket *model.Ticket) error {
	return nil
}

func (ts ticketImpl) GetTicketByID(id uint) (*model.Ticket, error) {
	return nil, nil
}
