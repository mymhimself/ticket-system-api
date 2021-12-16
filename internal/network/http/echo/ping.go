package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) Ping(ctx echo.Context) error {
	fmt.Println("salaam")
	return ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
}
