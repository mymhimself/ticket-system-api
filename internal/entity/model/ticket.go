package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
	"time"
)

type (
	Ticket struct {
		Thread   TicketThread
		Messages []TicketMessage
	}

	TicketThread struct {
		Title             string
		Priority          enum.Priority
		ClosedTime        time.Time
		Status            enum.TicketThreadStatus //
		CreatorID         uint
		Creator           User
		Department        int64
		AssignedToAdmin   *User
		AssignedToAdminID *uint
		gorm.Model
	}

	TicketMessage struct {
		ThreadID    uint
		Thread      TicketThread
		Text        string
		Seen        bool
		SentByAdmin bool
		Replied     bool
		ReplyTo     uint //self reference
		gorm.Model
	}
)
