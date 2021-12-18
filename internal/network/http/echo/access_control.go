package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/network/http/response"
	"net/http"
)

func accessControl(roles ...enum.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			_, _, role, err := extractUserInfoFromToken(ctx)
			if err != nil {
				return ctx.JSON(http.StatusUnauthorized, response.Message{
					Message: InvalidTokenErr,
					Status:  http.StatusUnauthorized,
				})
			}
			for _, r := range roles {
				if role == r {
					return next(ctx)
				}
			}

			return ctx.JSON(http.StatusUnauthorized, response.Message{
				Message: AccessForbiddenErr,
				Status:  http.StatusUnauthorized,
			})
		}
	}
}
