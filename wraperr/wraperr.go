package wraperr

import (
	"fmt"
	"runtime/debug"
)

type WrapError struct {
	Inner      error
	Message    string
	StackTrace string
}

func (err WrapError) Error() string {
	return err.Message
}

func Wrap(err error, message string, msgArgs ...interface{}) WrapError {
	return WrapError{
		Inner:      err,
		Message:    fmt.Sprintf(message, msgArgs...),
		StackTrace: string(debug.Stack()),
	}
}
