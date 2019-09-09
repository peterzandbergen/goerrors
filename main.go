package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testerror/typeerror"
	"testerror/varerror"
)

func main() {

	// Use wrapped error with error variables.
	var err error = varerror.ErrNotFound
	if errors.Is(err, varerror.ErrNotFound) {
		fmt.Println("one")
	}
	// Test if the error was because of sql.ErrNoRows.
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("two")
	}

	// Use wrapped errors with error types.
	err = typeerror.SystemFailure()
	var sferr *typeerror.ErrSystemFailure
	if errors.As(err, &sferr) {
		fmt.Println("three")
	}
	// Test if the error was because of sql.ErrNoRows.
	var sqlErr = sql.ErrNoRows
	if errors.Is(err, sqlErr) {
		fmt.Println("four")
	}
}
