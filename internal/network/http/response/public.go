package response

type (
	Message struct {
		Message string `json:"message"`
		Status  uint   `json:"status"`
	}
)
