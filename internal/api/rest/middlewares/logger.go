package middlewares

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
	"server/pkg/logger"
)

func registerLogger(restGroup *echo.Group) {
	restGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			echo_data.NewSetter(echoCTX).Logger(logger.New())
			return next(echoCTX)
		}
	})
}
