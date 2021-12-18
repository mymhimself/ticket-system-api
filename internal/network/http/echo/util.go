package echo

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func extractUserInfoFromToken(ctx echo.Context) (string, uint, enum.Role, error) {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return "", 0, 0, errors.New(InvalidTokenErr)
	}
	claims := token.Claims.(*model.TokenClaims)
	return claims.Username, claims.ID, claims.Role, nil
}
