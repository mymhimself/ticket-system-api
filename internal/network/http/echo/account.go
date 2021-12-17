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
	logger                logger.Logger
	accountService        service.Account
	ticketService         service.Ticket
	authenticationService service.Authentication
	validationService     service.Validation
}

func (h *handler) login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	if err := h.validationService.LoginRequest(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	token, refreshToken, err := h.accountService.Login(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	}

	return c.JSON(http.StatusOK, response.Login{RefreshToken: refreshToken, AccessToken: token, Status: http.StatusOK})
}

func (h *handler) register(c echo.Context) error {
	var req request.Register

	if err := c.Bind(&req); err != nil {
		h.logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	if err := h.validationService.RegisterRequest(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	}

	accessToken, refreshToken, err := h.accountService.Register(req.Username, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	}

	return c.JSON(http.StatusOK, response.Register{
		Message:      "User Successfully Registered",
		Status:       200,
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	})
}
