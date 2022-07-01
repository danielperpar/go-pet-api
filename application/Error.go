package application

type Error struct {
	Message string
}

func NewError(message string) *Error {
	return &Error{Message: message}
}