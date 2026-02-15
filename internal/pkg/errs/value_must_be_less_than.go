package errs

import (
	"errors"
	"fmt"
)

var ErrValueMustBeLessThan = errors.New("value must be less than maximum")

type ValueMustBeLessThanError struct {
	ParamName string
	Value     any
	Max       any
	Cause     error
}

func NewValueMustBeLessThan(paramName string, value, max any) error {
	return &ValueMustBeLessThanError{
		ParamName: paramName,
		Value:     value,
		Max:       max,
	}
}

func WrapValueMustBeLessThan(paramName string, value, max any, cause error) error {
	return &ValueMustBeLessThanError{
		ParamName: paramName,
		Value:     value,
		Max:       max,
		Cause:     cause,
	}
}

func (e *ValueMustBeLessThanError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(
			"%s: %s (%v) must be less than %v: %v",
			ErrValueMustBeLessThan,
			e.ParamName,
			e.Value,
			e.Max,
			e.Cause,
		)
	}

	return fmt.Sprintf(
		"%s: %s (%v) must be less than %v",
		ErrValueMustBeLessThan,
		e.ParamName,
		e.Value,
		e.Max,
	)
}

func (e *ValueMustBeLessThanError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueMustBeLessThan
}
