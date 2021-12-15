package model

import (
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"gorm.io/gorm"
)

type User struct {
	Username        string //email
	Phone           string
	PhoneIsVerified bool
	Password        string
	Role            enum.Role
	gorm.Model
}
