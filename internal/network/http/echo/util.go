package echo

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
)

func extractUserInfoFromToken(ctx echo.Context) (string, uint) {
	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.TokenClaims)
	return claims.Username, claims.ID
}
