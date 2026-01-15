package domain

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyExists   = errors.New("already exists")
	ErrEntityNotFound  = errors.New("entity not found")
	ErrIllegalArgument = errors.New("illegal Argument")
	ErrUnauthenticated = errors.New("no authentication found")
	ErrAccessDenied    = errors.New("access denied")
	ErrUserNotFound    = errors.New("user not found")
	ErrWrongPassword   = errors.New("wrong password")
)

func FormatError(e error, message string, args ...any) error {
	return fmt.Errorf(e.Error()+": "+message, args...)
}
