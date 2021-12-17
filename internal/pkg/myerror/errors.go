package myerror

import "github.com/mymhimself/ticket-system-api/internal/entity/enum"

type MyError struct {
	Code            ErrorCode  `json:"code"`      //e.g internal error
	SubCode         ErrorCode  `json:"sub_code""` //database connection failed
	Layer           enum.Layer `json:"layer"`     //repo, service, network
	ErrorMessage    string     `json:"error_message"`
	Trace           string     `json:"trace"`
	InternalMessage string     `json:"internal_message"`
}

func (e MyError) Error() string {
	return e.Code.String()
}

func New(code ErrorCode, layer enum.Layer, errorMessage string) *MyError {
	var err MyError
	err.Code = code
	err.InternalMessage = code.String()
	err.ErrorMessage = errorMessage
	err.Layer = layer
	return &err
}
