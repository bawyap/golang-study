package main

import "fmt"

type ValidationError struct {
	Field  string
	errMsg string
}

func (errr *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s: %s", errr.Field, errr.errMsg)
}
