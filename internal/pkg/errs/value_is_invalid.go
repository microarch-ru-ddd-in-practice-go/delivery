package errs

import (
	"errors"
	"fmt"
)

var ErrValueIsInvalid = errors.New("value is invalid")

type ValueIsInvalidError struct {
	ParamName string
	Cause     error
}

func NewValueIsInvalid(paramName string) error {
	return &ValueIsInvalidError{
		ParamName: paramName,
	}
}

func WrapValueIsInvalid(paramName string, cause error) error {
	return &ValueIsInvalidError{
		ParamName: paramName,
		Cause:     cause,
	}
}

func (e *ValueIsInvalidError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v",
			ErrValueIsInvalid,
			e.ParamName,
			e.Cause,
		)
	}

	return fmt.Sprintf("%s: %s",
		ErrValueIsInvalid,
		e.ParamName,
	)
}

func (e *ValueIsInvalidError) Unwrap() error {
	if e.Cause != nil {
		return e.Cause
	}
	return ErrValueIsInvalid
}
