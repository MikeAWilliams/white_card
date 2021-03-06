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

func (err WrapError) HasInnerError() bool {
	return nil != err.Inner
}

func Wrap(err error, message string, msgArgs ...interface{}) *WrapError {
	result := WrapError{
		Inner:      err,
		Message:    fmt.Sprintf(message, msgArgs...),
		StackTrace: string(debug.Stack()),
	}
	return &result
}

func Make(message string, msgArgs ...interface{}) *WrapError {
	result := WrapError{
		Inner:      nil,
		Message:    fmt.Sprintf(message, msgArgs...),
		StackTrace: string(debug.Stack()),
	}
	return &result
}
