package service

import "github.com/mymhimself/ticket-system-api/internal/network/http/request"

//validate request params
type Validation interface {
	LoginRequest(requestBody *request.Login) error
	NewTicketRequest(requestBody *request.NewTicket) error
	ReplyTicketRequest(requestBody *request.ReplyTicket) error
	RegisterConfirmRequest(requestBody *request.RegisterConfirm) error
	RegisterRequest(requestBody *request.Register) error
	ChangeTicketStatus(requestBody *request.ChangeTicketStatus) error
}
