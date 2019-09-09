package main

import (
	"database/sql"
	"errors"
	"fmt"
	"testerror/typeerror"
	"testerror/varerror"
)

func main() {
	var err error = varerror.ErrNotFound
	// Test errors.Is
	if errors.Is(err, varerror.ErrNotFound) {
		fmt.Println("one")
	}
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("two")
	}

	err = typeerror.SystemFailure(sql.ErrNoRows)
	var sferr *typeerror.ErrSystemFailure
	if errors.As(err, &sferr) {
		fmt.Println("three")
	}
	var sqlErr error = sql.ErrNoRows
	if errors.Is(err, sqlErr) {
		fmt.Println("four")
	}
}
