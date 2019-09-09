package typeerror

type myError struct {
	msg  string
	werr error
}

type ErrNotFound struct {
	myError
}

func newErrNotFound(err error) error {
	return &ErrNotFound{
		myError: myError{
			msg:  "error not found",
			werr: err,
		},
	}
}

func (e *myError) Error() string {
	return "oops"
}

func (e *myError) Unwrap() error {
	return e.werr
}

type ErrSystemFailure struct {
	myError
}

func newErrSystemFailure(err error) error {
	return &ErrSystemFailure{
		myError: myError{
			msg:  "system failure",
			werr: err,
		},
	}
}

func SystemFailure(err error) error {
	return newErrSystemFailure(err)
}

func NotFound(err error) error {
	return newErrNotFound(err)
}
