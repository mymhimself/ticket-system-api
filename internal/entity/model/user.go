package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
)

type User struct {
	Username       string //email
	HashedPassword []byte `gorm:"type:bytea"` //postgres type: bytea (byte array)
	Salt           string
	Role           enum.Role
	gorm.Model
}
