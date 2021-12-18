package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/network/http/request"
	"github.com/mymhimself/ticket-system-api/internal/network/http/response"
	"github.com/mymhimself/ticket-system-api/internal/service/ticket"
	"net/http"
	"strconv"
)

func (h *handler) createNewTicket(ctx echo.Context) error {
	var body request.NewTicket

	if err := ctx.Bind(&body); err != nil {
		fmt.Println(err.Error())
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest,
		})
	} else {
		//validation process
		if err := h.validationService.NewTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else {
			//create ticket and call CreateNewTicket
			var ticketThread model.TicketThread
			ticketThread.Priority = body.Priority
			ticketThread.Department = body.Department
			ticketThread.Title = body.Title
			ticketThread.Status = enum.Active
			_, ticketThread.CreatorID, _, _ = extractUserInfoFromToken(ctx)

			if err := h.ticketService.CreateNewTicket(&ticketThread, body.Text); err != nil {
				return ctx.JSON(http.StatusInternalServerError, response.Message{
					Message: err.Error(),
					Status:  http.StatusInternalServerError,
				})
			}
		}
		return ctx.JSON(
			http.StatusOK,
			response.Message{
				Message: TicketCreated,
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) replyToTicketMessage(ctx echo.Context) error {
	var body request.ReplyTicket

	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest},
		)
	} else {
		//validation process
		if err := h.validationService.ReplyTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else {
			_, senderID, _, _ := extractUserInfoFromToken(ctx)
			err := h.ticketService.ReplyMessage(body.Text, body.ReplyTo, senderID)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.Message{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				})
			}

		}
		return ctx.JSON(
			http.StatusOK,
			response.Message{
				Message: TicketMessageSaved,
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) replyToTicketMessageByAdmin(ctx echo.Context) error {
	var body request.ReplyTicket

	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest},
		)
	} else {
		//validation process
		if err := h.validationService.ReplyTicketRequest(&body); err != nil {
			return ctx.JSON(http.StatusBadRequest, response.Message{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
		} else {
			_, senderID, _, _ := extractUserInfoFromToken(ctx)
			err := h.ticketService.ReplyMessageByAdmin(body.Text, body.ReplyTo, senderID)
			if err != nil {
				return ctx.JSON(http.StatusBadRequest, response.Message{
					Message: err.Error(),
					Status:  http.StatusBadRequest,
				})
			}

		}
		return ctx.JSON(
			http.StatusOK,
			response.Message{
				Message: TicketMessageSaved,
				Status:  http.StatusOK,
			},
		)
	}
}

func (h *handler) ticketList(ctx echo.Context) error {
	var body request.FilterTicket
	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest},
		)
	}

	_, senderID, _, _ := extractUserInfoFromToken(ctx)
	list, err := h.ticketService.GetFilteredTicketThreads(
		&senderID,
		body.Seen,
		body.Status,
		body.Replied,
		body.Priority,
		body.Department,
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}

	return ctx.JSON(
		http.StatusOK,
		response.TicketList{
			Status: http.StatusOK,
			List:   list,
		},
	)
}

func (h *handler) changeTicketStatusByAdmin(ctx echo.Context) error {
	var body request.ChangeTicketStatus
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest},
		)
	}
	err := h.validationService.ChangeTicketStatus(&body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
			Status:  http.StatusBadRequest},
		)
	}

	err = h.ticketService.ChangeTicketStatus(body.TicketID, body.Status)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
			Status:  http.StatusBadRequest},
		)
	} else {
		return ctx.JSON(http.StatusOK, response.Message{
			Message: TicketStatusChanged,
			Status:  http.StatusOK,
		},
		)
	}
}

func (h *handler) ticketListByAdmin(ctx echo.Context) error {
	var body request.FilterTicket
	//body extraction process
	if err := ctx.Bind(&body); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest},
		)
	}

	//_, senderID := extractUserInfoFromToken(ctx)
	list, err := h.ticketService.GetFilteredTicketThreads(
		nil,
		body.Seen,
		body.Status,
		body.Replied,
		body.Priority,
		body.Department,
	)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
	}

	return ctx.JSON(
		http.StatusOK,
		response.TicketList{
			Status: http.StatusOK,
			List:   list,
		},
	)
}

func (h *handler) ticketDetails(ctx echo.Context) error {
	id := ctx.Param("id")
	ticketID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest,
		})
	}
	ticketDetail, exist, err := h.ticketService.GetTicketByID(uint(ticketID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	} else if !exist {
		return ctx.JSON(http.StatusOK, response.Message{
			Message: ticket.NotFoundTicket,
			Status:  http.StatusOK,
		})
	}

	return ctx.JSON(http.StatusOK, response.TicketInfo{
		Status: http.StatusOK,
		Ticket: *ticketDetail,
	})
}

func (h *handler) ticketDetailsByAdmin(ctx echo.Context) error {
	id := ctx.Param("id")
	ticketID, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.Message{
			Message: BadRequestFormat,
			Status:  http.StatusBadRequest,
		})
	}
	ticketDetail, exist, err := h.ticketService.GetTicketByID(uint(ticketID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Message{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	} else if !exist {
		return ctx.JSON(http.StatusOK, response.Message{
			Message: ticket.NotFoundTicket,
			Status:  http.StatusOK,
		})
	}
	for _, message := range ticketDetail.Messages {
		if !message.Seen {
			message.Seen = true
			h.ticketService.UpdateTicketMessage(&message)
		}
	}
	return ctx.JSON(http.StatusOK, response.TicketInfo{
		Status: http.StatusOK,
		Ticket: *ticketDetail,
	})
}
