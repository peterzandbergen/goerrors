package varerror

import (
	"database/sql"
	"fmt"
)

type myError struct {
	wrappedError error
	msg          string
}

// Using this approach does not allow us to have more than one cause
// for the error.

// ErrNotFound is defined once with the same wrapped error.
var ErrNotFound = newError("no result found", sql.ErrNoRows)

// ErrNotFound2 wraps another error.
var ErrNotFound2 = newError("no result found because of another error", sql.ErrConnDone)

// newError creates an error with a fixed wrapped error.
func newError(msg string, wrapped error) error {
	err := &myError{
		msg:          msg,
		wrappedError: wrapped,
	}
	return err
}

// Error returns the error cause.
func (e *myError) Error() string {
	return fmt.Sprintf("ErrNotFound, caused by internal error:\n %s", e.wrappedError.Error())
}

// Unwrap returns the wrapped error.
func (e *myError) Unwrap() error {
	return e.wrappedError
}
