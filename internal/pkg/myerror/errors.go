package myerror

import "github.com/mymhimself/ticket-system-api/internal/entity/enum"

type MyError struct {
	Code            ErrorCode  `json:"code"`
	Layer           enum.Layer `json:"layer"` //repo, service, network
	ErrorMessage    string     `json:"error_message"`
	InternalMessage string     `json:"internal_message"`
}

func (e MyError) Error() string {
	return e.Code.String()
}

func New(code ErrorCode, layer enum.Layer, errorMessage string) MyError {
	var err MyError
	err.Code = code
	err.InternalMessage = code.String()
	err.ErrorMessage = errorMessage
	err.Layer = layer
	return err
}
