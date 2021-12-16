package myerror

type ErrorCode uint

const (
	UserNotFound = iota
	InternalError
	PasswordIncorrect
	TicketNotFound
	InvalidPasswordLength
	UserExists
	UserSuccessfullyRegistered
)

var stringMap map[ErrorCode]string = map[ErrorCode]string{
	UserNotFound:               "User not found.",
	InternalError:              "Internal server error.",
	PasswordIncorrect:          "Incorrect password.",
	TicketNotFound:             "Ticket not found.",
	InvalidPasswordLength:      "Invalid Password length. it's length most be between 8 and 30.",
	UserExists:                 "This username is already taken.",
	UserSuccessfullyRegistered: "User successfully registered.",
}

func (e ErrorCode) String() string {
	return stringMap[e]
}
