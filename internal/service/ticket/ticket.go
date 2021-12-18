package ticket

import (
	"errors"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"time"
)

func (ts ticketImpl) CreateNewTicket(thread *model.TicketThread, text string) error {
	var message model.TicketMessage
	message.Text = text
	message.Thread = *thread
	message.ThreadID = thread.ID
	message.Seen = false
	message.SenderID = thread.CreatorID

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
		//ticket message not found
		return errors.New(MessageNotFound)
	}
	//another creator
	if repliedMessage.Thread.CreatorID != senderID {
		return errors.New(MessageNotFound)
	}
	//ticket is closed
	if repliedMessage.Thread.Status == enum.Closed {
		return errors.New(ClosedTicket)
	}

	var newMessage model.TicketMessage
	newMessage.SenderID = senderID
	newMessage.ReplyTo = repliedMessage.ID
	newMessage.ThreadID = repliedMessage.ThreadID
	newMessage.Text = text

	if err := ts.postgres.CreateMessageWithReply(&newMessage, repliedMessage); err != nil {
		return err
	}

	return nil
}

func (ts ticketImpl) ReplyMessageByAdmin(text string, repliedMessageID uint, senderID uint) error {
	repliedMessage, exist, err := ts.postgres.GetTicketMessageByID(repliedMessageID)
	if err != nil {
		return err
	} else if !exist {
		return errors.New(MessageNotFound)
	}

	var newMessage model.TicketMessage
	newMessage.SenderID = senderID
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

func (ts ticketImpl) UpdateTicketMessage(ticket *model.TicketMessage) error {
	return ts.postgres.UpdateMessage(ticket)
}

func (ts ticketImpl) GetTicketByID(id uint) (*model.Ticket, bool, error) {
	thread, exist, err := ts.postgres.GetTicketThreadByID(id)
	if err != nil {
		return nil, false, err
	} else if !exist {
		return nil, false, nil
	}
	messages, err := ts.postgres.GetTicketMessageByThreadID(thread.ID)
	if err != nil {
		return nil, false, err
	}
	var ticket model.Ticket
	ticket.Thread = *thread
	ticket.Messages = messages
	return &ticket, true, nil
}

func (ts ticketImpl) GetFilteredTicketThreads(creatorID *uint, seen *bool, status *enum.TicketThreadStatus, replied *bool, priority *enum.Priority, department *int64) ([]model.TicketInfo, error) {
	var filterMap = make(map[string]interface{})
	if seen != nil {
		filterMap["seen"] = *seen //
	}
	if replied != nil {
		filterMap["replied"] = *replied //
	}
	if status != nil {
		filterMap["status"] = *status
	}
	if priority != nil {
		filterMap["priority"] = *priority
	}
	if department != nil {
		filterMap["department"] = *department
	}
	list, err := ts.postgres.GetUserTicketThreadListByFilter(creatorID, filterMap)
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

func (ts ticketImpl) ChangeTicketStatus(ticketID uint, status enum.TicketThreadStatus) error {
	ticket, exist, err := ts.postgres.GetTicketThreadByID(ticketID)
	if err != nil {
		return err
	} else if !exist {
		return errors.New(NotFoundTicket)
	} else {
		ticket.Status = status
		if ticket.Status == enum.Closed {
			ticket.ClosedTime = time.Now()
		}
		return ts.postgres.UpdateTicketThread(ticket)
	}
}
