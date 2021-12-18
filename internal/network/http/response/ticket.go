package response

import "github.com/mymhimself/ticket-system-api/internal/entity/model"

type (
	TicketList struct {
		Status uint               `json:"status"`
		List   []model.TicketInfo `json:"list"`
	}

	TicketInfo struct {
		Status uint         `json:"status"`
		Ticket model.Ticket `json:"ticket"`
	}
)
