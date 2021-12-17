package request

type (
	NewTicket struct {
		Title      string `json:"title" validate:"required,len<80"`
		Text       string `json:"text" validate:"required,len<400"`
		Priority   int64  `json:"priority" validate:"required"`
		Department int64  `json:"department" validate:"required"`
	}

	ReplyTicket struct {
		Text    string `json:"text"`
		ReplyTo uint   `json:"reply_to"`
	}
)
