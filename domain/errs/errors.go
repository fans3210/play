package errs

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUnexpected   = errors.New("unexpected error")
)
