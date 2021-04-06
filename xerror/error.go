package xerror

import (
	"fmt"
)

type Error struct {
	PError  error
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	if e.PError == nil {
		return fmt.Sprintf("code: %s, message: %s", e.Code, e.Message)
	}

	return fmt.Sprintf("%s, code: %s, message: %s", e.PError, e.Code, e.Message)
}

func NewError(err error, code string, message string) *Error {
	return &Error{
		PError:  err,
		Code:    code,
		Message: message,
	}
}
