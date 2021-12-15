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
)
