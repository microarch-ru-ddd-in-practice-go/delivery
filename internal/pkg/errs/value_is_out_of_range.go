package errs

import (
	"errors"
	"fmt"
	"strings"
)

var ErrMustBeBetween = errors.New("must be between")

type MustBeBetweenError struct {
	ParamName string
	Value     any
	Min       any
	Max       any
	Cause     error
}

func NewMustBeBetween(paramName string, value, min, max any) error {
	return &MustBeBetweenError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
		Max:       max,
	}
}

func WrapMustBeBetween(paramName string, value, min, max any, cause error) error {
	return &MustBeBetweenError{
		ParamName: paramName,
		Value:     value,
		Min:       min,
		Max:       max,
		Cause:     cause,
	}
}

func (e *MustBeBetweenError) Error() string {
	value := sanitize(e.Value)

	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (%v) must be between %v and %v: %v",
			ErrMustBeBetween,
			e.ParamName,
			value,
			e.Min,
			e.Max,
			e.Cause,
		)
	}

	return fmt.Sprintf("%s: %s (%v) must be between %v and %v",
		ErrMustBeBetween,
		e.ParamName,
		value,
		e.Min,
		e.Max,
	)
}

func (e *MustBeBetweenError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrMustBeBetween
}

func sanitize(input any) string {
	str := fmt.Sprintf("%v", input)
	return strings.ReplaceAll(str, "\n", " ")
}
