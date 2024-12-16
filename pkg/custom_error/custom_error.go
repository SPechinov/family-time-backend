package custom_error

type CustomError struct {
	code string
}

func (customError *CustomError) Error() string {
	return customError.code
}

func (customError *CustomError) Code() string {
	return customError.code
}

func New(code string) *CustomError {
	return &CustomError{
		code: code,
	}
}
