package custom_errors

import "errors"

type ErrorBadRequest BaseError

func (err *ErrorBadRequest) Error() string {
	return err.Error()
}

func BadRequestError() error {
	return &ErrorBadRequest{
		Err:        errors.New("Bad Request"),
		StatusCode: 400,
		Message:    "Bad Request",
	}
}
