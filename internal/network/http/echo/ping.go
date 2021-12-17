package echo

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Ping(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
}
