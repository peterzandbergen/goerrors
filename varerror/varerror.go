package varerror

import (
	"database/sql"
	"fmt"
)

type myError struct {
	wrappedError error
	msg          string
}

var ErrNotFound = newError("no result found", sql.ErrNoRows)

func newError(msg string, wrapped error) error {
	err := &myError{
		msg:          msg,
		wrappedError: wrapped,
	}
	return err
}

func (e *myError) Error() string {
	return fmt.Sprintf("ErrNotFound, caused by internal error:\n %s", e.wrappedError.Error())
}

func (e *myError) Unwrap() error {
	return e.wrappedError
}
