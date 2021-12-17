package response

type (
	Login struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Status       int    `json:"status"`
	}

	Error struct {
		Error  string `json:"error"`
		Status int    `json:"status"`
	}

	Register struct {
		Message      string `json:"message"`
		Status       int    `json:"status"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
