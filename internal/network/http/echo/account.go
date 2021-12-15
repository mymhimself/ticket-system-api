package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/network/http/request"
	"github.com/mymhimself/ticket-system-api/internal/network/http/response"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/service"
	"net/http"
)

type handler struct {
	logger         logger.Logger
	accountService service.Account
	ticketService  service.Ticket
}

func (h *handler) login(c echo.Context) error {
	// start span with context

	req := new(request.Login)

	if err := c.Bind(req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	token, err := h.accountService.Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	}

	return c.JSON(http.StatusOK, response.Login{Token: token, Status: http.StatusOK})
}

func (h *handler) register(c echo.Context) error {
	return nil
}

func (h *handler) registerConfirm(c echo.Context) error {
	return nil
}
