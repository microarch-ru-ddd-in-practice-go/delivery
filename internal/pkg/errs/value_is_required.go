package errs

import (
	"errors"
	"fmt"
)

var ErrValueIsRequired = errors.New("value is required")

type ValueIsRequiredError struct {
	ParamName string
	Cause     error
}

func NewValueIsRequired(paramName string) error {
	return &ValueIsRequiredError{
		ParamName: paramName,
	}
}

func WrapValueIsRequired(paramName string, cause error) error {
	return &ValueIsRequiredError{
		ParamName: paramName,
		Cause:     cause,
	}
}

func (e *ValueIsRequiredError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v",
			ErrValueIsRequired,
			e.ParamName,
			e.Cause,
		)
	}

	return fmt.Sprintf("%s: %s",
		ErrValueIsRequired,
		e.ParamName,
	)
}

func (e *ValueIsRequiredError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueIsRequired
}
