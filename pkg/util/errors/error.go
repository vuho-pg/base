package errors

import "errors"

type ErrorType uint

const (
	ErrUnknown ErrorType = iota
	ErrBadRequest
	ErrInternal
)

type Error struct {
	Type ErrorType
	base error
	msg  string
}

func (e Error) Error() string {
	if e.msg == "" {
		return e.msg
	}
	return e.base.Error()
}

func (e Error) Unwrap() error {
	return e.base
}

func BadRequest(msg string) error {
	return Error{
		Type: ErrBadRequest,
		base: nil,
		msg:  msg,
	}
}

func Wrap(err error, t ErrorType) error {
	return Error{
		Type: t,
		base: err,
		msg:  "",
	}
}

func New(msg string) error {
	return Error{
		Type: ErrUnknown,
		base: nil,
		msg:  msg,
	}
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
