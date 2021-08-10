package verrors

import (
	"errors"
)

// Err general error
type Err struct {
	Cause error
	ID    string // optional
}

func (e Err) Error() string {
	var ret string
	if e.Cause != nil {
		ret = e.Cause.Error()
	}

	if e.ID != "" {
		ret = "ID(" + e.ID + "): " + ret
	}

	return ret
}

// TODO: add Unwrap(err error) error

type InvalidArgument struct {
	Err
}

func NewInvalidArgument(cause error) error {
	return InvalidArgument{Err{
		Cause: cause,
	}}
}

func IsInvalidArgument(err error) bool {
	var target InvalidArgument
	ok := errors.As(err, &target)

	return ok
}

type NotFound struct {
	Err
}

// NewNotFound returns NotFound error
func NewNotFound(cause error, id ...string) error {
	err := NotFound{Err{Cause: cause}}

	if len(id) > 0 {
		err.ID = id[0]
	}

	return err
}

func IsNotFound(err error) bool {
	var target NotFound
	ok := errors.As(err, &target)

	return ok
}

func (e NotFound) Error() string {
	if e.Cause == nil {
		e.Cause = errors.New("not found")
	}

	return e.Err.Error()
}

type AlreadyExists struct {
	Err
}

func NewAlreadyExists(cause error, id ...string) error {
	err := AlreadyExists{Err{Cause: cause}}

	if len(id) > 0 {
		err.ID = id[0]
	}

	return err
}

func IsAlreadyExists(err error) bool {
	var target AlreadyExists
	ok := errors.As(err, &target)

	return ok
}

func (e AlreadyExists) Error() string {
	if e.Cause == nil {
		e.Cause = errors.New("already exists")
	}

	return e.Err.Error()
}
