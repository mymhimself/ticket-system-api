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

	TicketInfo struct {
		ID         uint                    `json:"id"`
		SenderID   uint                    `json:"sender_id"`
		Title      string                  `json:"title"`
		Priority   enum.Priority           `json:"priority"`
		Department uint                    `json:"department"`
		Status     enum.TicketThreadStatus `json:"status"`
		Seen       bool                    `json:"seen"`
		Replied    bool                    `json:"replied"`
	}

	//database table
	TicketThread struct {
		Title             string
		Priority          enum.Priority
		ClosedTime        time.Time
		Status            enum.TicketThreadStatus //
		CreatorID         uint
		Creator           User `gorm:"foreignKey:CreatorID"`
		Department        int64
		AssignedToAdmin   *User
		AssignedToAdminID *uint
		gorm.Model
	}

	//database table
	TicketMessage struct {
		ThreadID uint
		Thread   TicketThread `gorm:"foreignKey:ThreadID"`
		Text     string
		Seen     bool
		SenderID uint
		Sender   User `gorm:"foreignKey:SenderID"`
		Replied  bool
		ReplyTo  uint //self reference
		gorm.Model
	}
)
