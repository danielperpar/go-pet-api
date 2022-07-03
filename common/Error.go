package common

import ("fmt")

//TODO: quizás no necesarios, revisar
const (
	Domain_NoPets       	string = "There are no pets"
	Infrast_DbErrorOpen		string = "Open db connection Error"
	Infrast_DbErrorInsert	string = "Database error on Insert"
	Infrast_DbErrorQuery	string = "Database error on Query"
	Infrast_DbErrorUnknown  string = "Unknown database error"
	Infrast_DbQueryResultScanError string = "Error on scanning query result"
	UnknownError 	string = "Unknown error"
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
