package rest_error

type RestError struct {
	HttpCode int
	Code     string
}

func (e *RestError) Error() string {
	return e.Code
}

func New(httpCode int, code string) *RestError {
	return &RestError{
		HttpCode: httpCode,
		Code:     code,
	}
}

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewValidation(message string) error {
	return &ValidationError{
		Message: message,
	}
}
