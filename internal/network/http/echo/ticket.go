package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/network/http/request"
	"github.com/mymhimself/ticket-system-api/internal/network/http/response"
	"net/http"
)

func (h *handler) createNewTicket(ctx echo.Context) error {
	var body request.NewTicket

	if err := ctx.Bind(&body); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Error{
			Error:  "bad request format",
			Status: http.StatusBadRequest,
		})
	} else {
		//validation process
		if err := h.validationService.NewTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Error{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})
		} else {
			//create ticket and call CreateNewTicket
			var ticketThread model.TicketThread
			ticketThread.Priority = enum.Priority(body.Priority)
			ticketThread.Department = body.Department
			ticketThread.Title = body.Title
			ticketThread.Status = enum.Active
			_, ticketThread.CreatorID = extractUserInfoFromToken(ctx)

			if err := h.ticketService.CreateNewTicket(&ticketThread, body.Text); err != nil {
				return ctx.JSON(http.StatusInternalServerError, response.Error{
					Error:  err.Error(),
					Status: http.StatusInternalServerError,
				})
			}
		}
		return ctx.JSON(
			http.StatusOK,
			response.NewTicket{
				Message: "Ticket created and saved.",
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) replyToTicketMessage(ctx echo.Context) error {
	var body request.ReplyTicket

	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Error{
			Error:  "bad request format",
			Status: http.StatusBadRequest},
		)
	} else {
		//validation process
		if err := h.validationService.ReplyTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Error{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})
		} else {
			_, senderID := extractUserInfoFromToken(ctx)
			err := h.ticketService.ReplyMessage(body.Text, body.ReplyTo, senderID)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.Error{
					Error:  err.Error(),
					Status: http.StatusBadRequest,
				})
			}

		}
		return ctx.JSON(
			http.StatusOK,
			response.NewTicket{
				Message: "Message saved.",
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) replyToTicketMessageAsAdmin(ctx echo.Context) error {
	var body request.ReplyTicket

	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Error{
			Error:  "bad request format",
			Status: http.StatusBadRequest},
		)
	} else {
		//validation process
		if err := h.validationService.ReplyTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Error{
				Error:  err.Error(),
				Status: http.StatusBadRequest,
			})
		} else {
			_, senderID := extractUserInfoFromToken(ctx)
			err := h.ticketService.ReplyMessageAsAdmin(body.Text, body.ReplyTo, senderID)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.Error{
					Error:  err.Error(),
					Status: http.StatusBadRequest,
				})
			}

		}
		return ctx.JSON(
			http.StatusOK,
			response.NewTicket{
				Message: "Message saved.",
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) ticketList(ctx echo.Context) error {
	var body request.FilterTicket
	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Error{
			Error:  "bad request format",
			Status: http.StatusBadRequest},
		)
	}

	if body.Seen != nil {
		fmt.Println(*body.Seen)
	}

	_, senderID := extractUserInfoFromToken(ctx)
	list, err := h.ticketService.GetFilteredTicketThreads(
		senderID,
		body.Seen,
		body.Status,
		body.Replied,
		body.Priority,
		body.Department,
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Error{
			Error:  err.Error(),
			Status: http.StatusBadRequest,
		})
	}

	return ctx.JSON(
		http.StatusOK,
		response.TicketInfo{
			Status: http.StatusOK,
			List:   list,
		},
	)
}
