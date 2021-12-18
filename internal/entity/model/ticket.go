package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
	"time"
)

type (
	//not save in database
	Ticket struct {
		Thread   TicketThread    `json:"thread"`
		Messages []TicketMessage `json:"messages"`
	}

	//not save in database
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
		ID         uint                    `gorm:"primarykey" json:"id"`
		Title      string                  `json:"title"`
		Priority   enum.Priority           `json:"priority"`
		ClosedTime time.Time               `json:"closed_time,omitempty"`
		Status     enum.TicketThreadStatus `json:"status"` //
		CreatorID  uint                    `json:"creator_id"`
		Creator    User                    `gorm:"foreignKey:CreatorID" json:"-"`
		Department int64                   `json:"department"`
		//AssignedToAdmin   *User `json:"-"`
		//AssignedToAdminID *uint `json:"assigned_to_admin_id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	}

	//database table
	//gorm.Model `json:"model"`
	TicketMessage struct {
		ID        uint           `gorm:"primarykey" json:"id"`
		ThreadID  uint           `json:"-"`
		Thread    TicketThread   `gorm:"foreignKey:ThreadID" json:"-"`
		Text      string         `json:"text"`
		Seen      bool           `json:"seen"`
		SenderID  uint           `json:"sender_id"`
		Sender    User           `gorm:"foreignKey:SenderID" json:"-"`
		Replied   bool           `json:"replied"`
		ReplyTo   uint           `json:"reply_to"` //self reference
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	}
)
