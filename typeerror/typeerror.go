package typeerror

import "database/sql"

type myError struct {
	msg  string
	werr error
}

func (e *myError) Error() string { return e.msg }
func (e *myError) Unwrap() error { return e.werr }

// ErrNotFound is returned when no result was found.
// It can wrap any error that caused it.
type ErrNotFound struct{ myError }

// ErrSystemFailure is returned when no result was found.
// It can wrap any error that caused it.
type ErrSystemFailure struct{ myError }

// newErrNotFound creates returns a new ErrNotFound with wrapped error.
func newErrNotFound(msg string, err error) error {
	return &ErrNotFound{
		myError: myError{
			msg:  msg,
			werr: err,
		},
	}
}

// newErrSystemFailure allows wrapping more
func newErrSystemFailure(msg string, err error) error {
	return &ErrSystemFailure{
		myError: myError{
			msg:  msg,
			werr: err,
		},
	}
}

// SystemFailure does something and returns a system failure with a sql error wrapped.
func SystemFailure() error {
	err := sql.ErrNoRows
	return newErrSystemFailure("system failure", err)
}

// NotFound returns not found error with sql error wrapped.
func NotFound() error {
	err := sql.ErrNoRows
	return newErrNotFound("nothing found", err)
}
