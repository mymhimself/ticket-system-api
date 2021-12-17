package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
)

type (
	TokenClaims struct {
		ID       uint      `json:"id"`
		Username string    `json:"username"`
		Role     enum.Role `json:"role"`
		jwt.StandardClaims
	}

	RefreshTokenClaims struct {
		ID uint `json:"id"`
		jwt.StandardClaims
	}
)
