package errs

import (
	"errors"
	"fmt"
)

var ErrValueMustBeLessOrEqual = errors.New("value must be less than or equal to maximum")

type ValueMustBeLessOrEqualError struct {
	ParamName string
	Value     any
	Max       any
	Cause     error
}

func NewValueMustBeLessOrEqual(paramName string, value, max any) error {
	return &ValueMustBeLessOrEqualError{
		ParamName: paramName,
		Value:     value,
		Max:       max,
	}
}

func WrapValueMustBeLessOrEqual(paramName string, value, max any, cause error) error {
	return &ValueMustBeLessOrEqualError{
		ParamName: paramName,
		Value:     value,
		Max:       max,
		Cause:     cause,
	}
}

func (e *ValueMustBeLessOrEqualError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(
			"%s: %s (%v) must be less than or equal to %v: %v",
			ErrValueMustBeLessOrEqual,
			e.ParamName,
			e.Value,
			e.Max,
			e.Cause,
		)
	}

	return fmt.Sprintf(
		"%s: %s (%v) must be less than or equal to %v",
		ErrValueMustBeLessOrEqual,
		e.ParamName,
		e.Value,
		e.Max,
	)
}

func (e *ValueMustBeLessOrEqualError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueMustBeLessOrEqual
}
