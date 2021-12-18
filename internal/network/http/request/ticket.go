package request

import "github.com/mymhimself/ticket-system-api/internal/entity/enum"

type (
	NewTicket struct {
		Title      string        `json:"title"`
		Text       string        `json:"text"`
		Priority   enum.Priority `json:"priority"`
		Department int64         `json:"department"`
	}

	ReplyTicket struct {
		Text    string `json:"text"`
		ReplyTo uint   `json:"reply_to"`
	}

	FilterTicket struct {
		Seen    *bool `param:"seen" query:"seen" form:"seen"`
		Replied *bool `param:"replied" query:"replied"`

		Status     *enum.TicketThreadStatus `param:"status" query:"status"` //0=active, closed, suspend
		Department *int64                   `param:"department" query:"department"`
		Priority   *enum.Priority           `param:"priority" query:"priority"`
	}

	ChangeTicketStatus struct {
		Status   enum.TicketThreadStatus `json:"status"`
		TicketID uint                    `json:"ticket_id"`
	}
)
