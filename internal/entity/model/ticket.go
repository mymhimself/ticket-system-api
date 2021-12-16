package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
)

type (
	Ticket struct {
		Sender          User
		SenderID        uint
		Status          enum.TicketStatus
		Text            string
		Title           string
		Priority        enum.Priority
		ReplyToTicketID uint
		gorm.Model
	}
)
