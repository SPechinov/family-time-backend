package validate

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

var EmailErrMustBeString = validation.NewError("validation_email", "email must be a string")
var EmailErrEmptyString = validation.NewError("validation_phone", "empty string")
var EmailErrInvalid = validation.NewError("validation_email", "invalid email format")

func Email(value any) error {
	email, ok := value.(string)
	if !ok {
		return EmailErrMustBeString
	}

	if email == "" {
		return EmailErrEmptyString
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+[a-zA-Z0-9]@[a-zA-Z0-9]+[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return EmailErrInvalid
	}

	return nil
}
