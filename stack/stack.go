package stack

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type StackError struct {
	Err   error
	Stack string
}

func Wrap(err error) error {
	return wrap(debug.Stack, err)
}

func wrap(getStack func() []byte, err error) error {
	if err == nil {
		return nil
	}

	// Don't add an additional stack if the error already has one
	var stackErr *StackError
	if ok := errors.As(err, &stackErr); ok {
		return err
	}

	return &StackError{
		Err:   err,
		Stack: string(getStack()),
	}
}

func (e *StackError) Error() string {
	return fmt.Sprintf("%s\n%s", e.Err, e.Stack)
}

func (e *StackError) Unwrap() error {
	return e.Err
}
