package validate

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nyaruka/phonenumbers"
)

var PhoneErrMustBeString = validation.NewError("validation_phone", "phone must be a string")
var PhoneErrEmptyString = validation.NewError("validation_phone", "empty string")
var PhoneErrParse = validation.NewError("validation_phone", "invalid phone number")
var PhoneErrInvalid = validation.NewError("validation_phone", "invalid phone number")

func Phone(value any) error {
	phone, ok := value.(string)
	if !ok {
		return PhoneErrMustBeString
	}

	if phone == "" {
		return PhoneErrEmptyString
	}

	phoneNumber, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return PhoneErrParse
	}

	isValid := phonenumbers.IsValidNumber(phoneNumber)
	if !isValid {
		return PhoneErrInvalid
	}

	return nil
}
