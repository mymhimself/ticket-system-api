package request

type (
	NewTicket struct {
		Title      string `json:"title" validate:"required,len<80"`
		Text       string `json:"text" validate:"required,len<400"`
		Priority   uint   `json:"priority" validate:"required"`
		Department uint   `json:"department" validate:"required"`
	}
)
