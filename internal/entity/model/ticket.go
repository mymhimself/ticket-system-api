package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
)

type Ticket struct {
	Sender   User
	Text     string
	Title    string
	Priority enum.Priority
	gorm.Model
}
