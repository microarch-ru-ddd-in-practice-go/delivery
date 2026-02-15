package errs

import (
	"errors"
	"fmt"
)

var ErrValueMustBeGreaterOrEqual = errors.New("value must be greater than or equal to minimum")

type ValueMustBeGreaterOrEqualError struct {
	ParamName string
	Value     any
	Min       any
	Cause     error
}

func NewValueMustBeGreaterOrEqual(paramName string, value, min any) error {
	return &ValueMustBeGreaterOrEqualError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
	}
}

func WrapValueMustBeGreaterOrEqual(paramName string, value, min any, cause error) error {
	return &ValueMustBeGreaterOrEqualError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
		Cause:     cause,
	}
}

func (e *ValueMustBeGreaterOrEqualError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(
			"%s: %s (%v) must be greater than or equal to %v: %v",
			ErrValueMustBeGreaterOrEqual,
			e.ParamName,
			e.Value,
			e.Min,
			e.Cause,
		)
	}

	return fmt.Sprintf(
		"%s: %s (%v) must be greater than or equal to %v",
		ErrValueMustBeGreaterOrEqual,
		e.ParamName,
		e.Value,
		e.Min,
	)
}

func (e *ValueMustBeGreaterOrEqualError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueMustBeGreaterOrEqual
}
