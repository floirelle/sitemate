package custom_errors

type BaseError struct {
	StatusCode int
	Err        error
	Message    string `json:"message"`
}

func (e *BaseError) Error() string {
	return "base_error"
}
