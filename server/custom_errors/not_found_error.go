package custom_errors

import "errors"

type ErrNotFound BaseError

func (err *ErrNotFound) Error() string {
	return err.Err.Error()
}

func NotFoundError() error {
	return &BaseError{
		Err:        errors.New("Not found error"),
		StatusCode: 404,
		Message:    "Not found",
	}
}
