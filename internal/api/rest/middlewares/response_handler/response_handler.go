package mw_response_handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/api/rest/utils"
	"server/internal/api/rest/utils/echo_data"
	"server/internal/api/rest/utils/rest_error"
	"server/pkg/app_error"
	"server/pkg/custom_error"
)

func Register(restGroup *echo.Group) {
	restGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			err := next(echoCTX)
			if err == nil {
				return nil
			}

			var restError *rest_error.RestError
			var validationError *rest_error.ValidationError
			var customError *custom_error.CustomError
			var appError *app_error.AppError

			switch {
			case errors.As(err, &restError):
				return handleRestError(echoCTX, restError)
			case errors.As(err, &validationError):
				return handleValidationError(echoCTX, validationError)
			case errors.As(err, &customError):
				return handleCustomError(echoCTX, customError)
			case errors.As(err, &appError):
				return handlerAppError(echoCTX, appError)
			default:
				return handleDefault(echoCTX, err)
			}
		}
	})
}

func handleRestError(echoCTX echo.Context, restError *rest_error.RestError) error {
	return echoCTX.JSON(restError.HttpCode, utils.NewResponseBad(restError.Code))
}

func handleValidationError(echoCTX echo.Context, validationError *rest_error.ValidationError) error {
	return echoCTX.JSON(http.StatusBadRequest, utils.NewResponseBadValidation(validationError.Message))
}

func handlerAppError(echoCTX echo.Context, appError *app_error.AppError) error {
	l, _ := echo_data.NewGetter(echoCTX).Logger()
	l.WithStackTrace(appError.StackTrace())
	l.Error(appError)

	return echoCTX.JSON(
		rest_error.ErrSomethingHappened.HttpCode,
		utils.NewResponseBad(rest_error.ErrSomethingHappened.Code),
	)
}

func handleDefault(echoCTX echo.Context, err error) error {
	l, _ := echo_data.NewGetter(echoCTX).Logger()
	l.Error(err)

	return echoCTX.JSON(
		rest_error.ErrSomethingHappened.HttpCode,
		utils.NewResponseBad(rest_error.ErrSomethingHappened.Code),
	)
}
