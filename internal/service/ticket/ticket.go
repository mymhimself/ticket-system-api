package ticket

import (
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
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
	}
	//another creator
	if repliedMessage.Thread.CreatorID != senderID {
		return myerror.New(myerror.TicketNotFound, enum.ServiceLayer, "")
	}
	//ticket is closed
	if repliedMessage.Thread.Status == enum.Closed {
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
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

func (ts ticketImpl) ReplyMessageAsAdmin(text string, repliedMessageID uint, senderID uint) error {
	repliedMessage, exist, err := ts.postgres.GetTicketMessageByID(repliedMessageID)
	if err != nil {
		return err
	} else if !exist {
		return myerror.New(myerror.TicketMessageNotFound, enum.ServiceLayer, "")
	}
	if repliedMessage.Thread.AssignedToAdminID != &senderID {
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

func (ts ticketImpl) UpdateTicket(ticket *model.Ticket) error {
	return nil
}

func (ts ticketImpl) GetTicketByID(id uint) (*model.Ticket, error) {
	return nil, nil
}

func (ts ticketImpl) GetFilteredTicketThreads(creatorID uint, seen *bool, status *int64, replied *bool, priority *int64, department *int64) ([]model.TicketInfo, error) {
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
