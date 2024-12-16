package middlewares

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
)

func registerLogRequestID(restGroup *echo.Group) {
	restGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			echoGetter := echo_data.NewGetter(echoCTX)

			l, err := echoGetter.Logger()
			if err != nil {
				return err
			}

			requestID, err := echoGetter.RequestID()
			if err != nil {
				return err
			}

			l = l.WithRequestID(requestID)
			return next(echoCTX)
		}
	})
}
