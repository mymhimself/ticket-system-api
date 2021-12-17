package response

import "github.com/mymhimself/ticket-system-api/internal/entity/model"

type (
	NewTicket struct {
		Message string `json:"message"`
		Status  uint   `json:"status"`
	}

	TicketInfo struct {
		Status uint               `json:"status"`
		List   []model.TicketInfo `json:"list"`
	}
)
