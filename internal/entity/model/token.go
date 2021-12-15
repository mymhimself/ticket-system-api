package model

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
