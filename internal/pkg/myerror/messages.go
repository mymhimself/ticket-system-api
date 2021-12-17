package myerror

type ErrorCode uint

const (
	UserNotFound = iota
	InternalError
	PasswordIncorrect
	TicketMessageNotFound
	TicketIsClosed
	TicketNotFound
	InvalidPasswordLength
	UserExists
	UserSuccessfullyRegistered
	InvalidPriorityValue
	InvalidDepartmentValue
)

var stringMap map[ErrorCode]string = map[ErrorCode]string{
	UserNotFound:               "user not found",
	InternalError:              "internal server error",
	PasswordIncorrect:          "incorrect password",
	TicketMessageNotFound:      "ticket message not found",
	TicketIsClosed:             "ticket is closed",
	TicketNotFound:             "ticket not found",
	InvalidPasswordLength:      "invalid Password length. it's length most be between 8 and 30",
	UserExists:                 "this username is already taken",
	UserSuccessfullyRegistered: "user successfully registered",
	InvalidPriorityValue:       "invalid priority value",
	InvalidDepartmentValue:     "invalid department value",
}

func (e ErrorCode) String() string {
	return stringMap[e]
}
