package common

import (
	"fmt"
)

const (
	Domain_PetNotFound		string = "Species not found"
)

type Error struct {
	Code int
	Message string
}

func NewError(code int, message string) *Error{
	return &Error{Code: code, Message: message}
}

func (e *Error) Error() string {
	return fmt.Sprint("code: %d message: %s", e.Code, e.Message)
}
