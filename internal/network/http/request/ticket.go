package request

type (
	NewTicket struct {
		Title      string `json:"title"`
		Text       string `json:"text"`
		Priority   int64  `json:"priority"`
		Department int64  `json:"department"`
	}

	ReplyTicket struct {
		Text    string `json:"text"`
		ReplyTo uint   `json:"reply_to"`
	}

	FilterTicket struct {
		Seen    *bool `param:"seen" query:"seen" form:"seen"`
		Replied *bool `param:"replied" query:"replied"`

		Status     *int64 `param:"status" query:"status"` //0=active, 2=closed, 1=suspend
		Department *int64 `param:"department" query:"department"`
		Priority   *int64 `param:"priority" query:"priority"`
	}
)
