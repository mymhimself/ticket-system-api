package request

type (
	Login struct {
		Username string `json:"username"` //email
		Password string `json:"password"`
	}

	Register struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}

	RegisterConfirm struct {
		OTP      string `json:"otp" validate:"required"`
		Username string `json:"username" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Key      string `json:"key" validate:"required"`
	}
)
