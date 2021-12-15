package request

type (
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Register struct {
		Phone    string `json:"phone"`
		Username string `json:"username"`
	}

	RegisterConfirm struct {
	}
)
