package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/network/http/request"
	"net/http"
)

func (h *handler) newTicket(ctx echo.Context) error {
	var ticketBody request.NewTicket
	if err := ctx.Bind(&ticketBody); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	} else {
		if err := h.validationService.NewTicketRequest(&ticketBody); err != nil {
			return err
		} else {
			var ticket model.Ticket
			ticket.Priority = enum.Priority(ticketBody.Priority)
			ticket.Title = ticketBody.Title
			ticket.Text = ticketBody.Text
			ticket.Sender = ctx.Get("user").(model.User)
			ticket.SenderID = 1

			//h.ticketService.SendTicketByUser()
		}
		return nil
	}
}
