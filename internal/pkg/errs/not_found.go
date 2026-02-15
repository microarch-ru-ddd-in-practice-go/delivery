package errs

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

type NotFoundError struct {
	Name  string
	ID    any
	Cause error
}

func NewNotFound(name string, id any) error {
	return &NotFoundError{
		Name: name,
		ID:   id,
	}
}

func WrapNotFound(name string, id any, cause error) error {
	return &NotFoundError{
		Name:  name,
		ID:    id,
		Cause: cause,
	}
}

func (e *NotFoundError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (%v): %v",
			ErrNotFound,
			e.Name,
			e.ID,
			e.Cause,
		)
	}

	return fmt.Sprintf("%s: %s (%v)",
		ErrNotFound,
		e.Name,
		e.ID,
	)
}

func (e *NotFoundError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrNotFound
}
