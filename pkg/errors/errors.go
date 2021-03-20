package errors

import (
	"errors"
	"fmt"
)

var (
	New    = errors.New
	Is     = errors.Is
	As     = errors.As
	Unwrap = errors.Unwrap
)

func Newf(format string, args ...interface{}) error {
	text := fmt.Sprintf(format, args...)
	return New(text)
}

func Wrap(err error, text string) error {
	return fmt.Errorf("%s: %w", text, err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	text := fmt.Sprintf(format, args...)
	return Wrap(err, text)
}
