package response

type (
	NewTicket struct {
		Message string `json:"message"`
		Status  uint
	}
)
