package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
)

func registerRequestID(restGroup *echo.Group) {
	restGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			createdRequestID := uuid.New().String()

			echo_data.NewSetter(echoCTX).RequestID(createdRequestID)

			return next(echoCTX)
		}
	})
}
