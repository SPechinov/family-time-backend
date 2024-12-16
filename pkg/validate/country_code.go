package validate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"server/pkg/country"
)

var CountryCodeErrInvalid = validation.NewError("validation_country_code", "invalid country code")

func CountryCode(value any) error {
	countryCode, ok := value.(string)
	if !ok {
		return CountryCodeErrInvalid
	}

	if !country_codes.NewCodes().IsReal(countryCode) {
		return CountryCodeErrInvalid
	}

	return nil
}
