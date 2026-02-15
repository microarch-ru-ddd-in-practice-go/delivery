package errs

import (
	"errors"
	"fmt"
)

var ErrValueMustBeGreaterThan = errors.New("value must be greater than minimum")

type ValueMustBeGreaterThanError struct {
	ParamName string
	Value     any
	Min       any
	Cause     error
}

func NewValueMustBeGreaterThan(paramName string, value, min any) error {
	return &ValueMustBeGreaterThanError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
	}
}

func WrapValueMustBeGreaterThan(paramName string, value, min any, cause error) error {
	return &ValueMustBeGreaterThanError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
		Cause:     cause,
	}
}

func (e *ValueMustBeGreaterThanError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(
			"%s: %s (%v) must be greater than %v: %v",
			ErrValueMustBeGreaterThan,
			e.ParamName,
			e.Value,
			e.Min,
			e.Cause,
		)
	}

	return fmt.Sprintf(
		"%s: %s (%v) must be greater than %v",
		ErrValueMustBeGreaterThan,
		e.ParamName,
		e.Value,
		e.Min,
	)
}

func (e *ValueMustBeGreaterThanError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueMustBeGreaterThan
}
