package errors

import (
	"errors"
	"fmt"
)

func New(text string) error {
	return errors.New(text)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

func Wrap(op string, err error) error {
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
