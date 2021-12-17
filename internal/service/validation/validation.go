package validation

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/mymhimself/ticket-system-api/internal/network/http/request"
	passwordValidator "github.com/wagslane/go-password-validator"
)

func (s *validationImpl) LoginRequest(requestBody *request.Login) error {
	err := validation.ValidateStruct(requestBody,
		validation.Field(&requestBody.Username, validation.Required),
		validation.Field(&requestBody.Password, validation.Required),
	)

	return err
}

func (s *validationImpl) RegisterRequest(requestBody *request.Register) error {
	err := validation.ValidateStruct(requestBody,
		validation.Field(&requestBody.Username, validation.Required, validation.Length(5, 30), is.Email),
		validation.Field(&requestBody.Password, validation.Required, validation.Length(5, 30)),
	)
	if err != nil {
		return err
	}
	err = passwordValidator.Validate(requestBody.Password, 60)
	if err != nil {
		return err
	}
	return err
}

func (s *validationImpl) NewTicketRequest(requestBody *request.NewTicket) error {
	err := validation.ValidateStruct(requestBody,
		validation.Field(&requestBody.Text, validation.Required, validation.Length(10, 400)),
		validation.Field(&requestBody.Title, validation.Required, validation.Length(10, 80)),
		validation.Field(&requestBody.Priority, validation.Required, validation.Max(4), validation.Min(0)),
		validation.Field(&requestBody.Department, validation.Required, validation.Min(1)),
	)
	return err
}

//func (s *validationImpl) FilterTicketRequest(requestBody *request.FilterTicket) err {
//	err := validation.ValidateStruct(requestBody,
//		validation.Field(&requestBody.Text, validation.Required, validation.Length(10, 400)),
//		validation.Field(&requestBody.Title, validation.Required, validation.Length(10, 80)),
//		validation.Field(&requestBody.Priority, validation.Required, validation.Max(4), validation.Min(0)),
//		validation.Field(&requestBody.Department, validation.Required, validation.Min(1)),
//	)
//	return err
//}

func (s *validationImpl) ReplyTicketRequest(requestBody *request.ReplyTicket) error {
	err := validation.ValidateStruct(requestBody,
		validation.Field(&requestBody.Text, validation.Required, validation.Length(10, 400)),
		validation.Field(&requestBody.ReplyTo, validation.Required),
	)

	return err
}

func (s *validationImpl) RegisterConfirmRequest(requestBody *request.RegisterConfirm) error {
	err := validation.ValidateStruct(requestBody,
		validation.Field(&requestBody.Password, validation.Required, validation.Length(8, 30)),
		validation.Field(&requestBody.Username, validation.Required, validation.Length(5, 30), is.Email),
		validation.Field(&requestBody.Key, validation.Required),
		validation.Field(&requestBody.OTP, validation.Required),
	)
	if err != nil {
		return err
	}
	err = passwordValidator.Validate(requestBody.Password, 60)
	if err != nil {
		return err
	}
	return nil
}
