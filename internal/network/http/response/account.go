package response

type (
	Login struct {
		Token  string `json:"token"`
		Status int    `json:"status"`
	}

	Error struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}

	Register struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
)
