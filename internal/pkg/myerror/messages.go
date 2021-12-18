package myerror

type ErrorCode uint

const (
	_ = iota
	InternalError
)

var stringMap map[ErrorCode]string = map[ErrorCode]string{
	InternalError: "internal server error",
}

func (e ErrorCode) String() string {
	return stringMap[e]
}
