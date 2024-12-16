package mw_response_handler

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils"
	"server/internal/api/rest/utils/rest_error"
	"server/pkg/custom_error"
)

var appErrorMapping = map[string]*rest_error.RestError{
	custom_error.ErrCodeIsNotInRedis.Code():  rest_error.ErrCodeDidNotSent,
	custom_error.ErrCodesNotEqual.Code():     rest_error.ErrCodesNotEqual,
	custom_error.ErrCodeMaxAttempts.Code():   rest_error.ErrCodeMaxAttempts,
	custom_error.ErrUserExists.Code():        rest_error.ErrUserExists,
	custom_error.ErrUserNotExists.Code():     rest_error.ErrUserNotExists,
	custom_error.ErrUnknownAuthMethod.Code(): rest_error.ErrSomethingHappened,
	custom_error.ErrIncorrectPassword.Code(): rest_error.ErrIncorrectPassword,
	custom_error.ErrNotAuthorized.Code():     rest_error.ErrNotAuthorized,
}

func handleCustomError(echoCTX echo.Context, customError *custom_error.CustomError) error {
	value, exist := appErrorMapping[customError.Code()]
	if !exist {
		return echoCTX.JSON(
			rest_error.ErrSomethingHappened.HttpCode,
			utils.NewResponseBad(rest_error.ErrSomethingHappened.Code),
		)
	}

	return echoCTX.JSON(
		value.HttpCode,
		utils.NewResponseBad(value.Code),
	)
}
