package middlewares

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
)

func registerLogURI(restGroup *echo.Group) {
	restGroup.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			l, err := echo_data.NewGetter(echoCTX).Logger()
			if err != nil {
				return err
			}

			l = l.WithURI(echoCTX.Request().RequestURI)

			l.Debug("Start")
			err = next(echoCTX)
			l.Debug("Finish")

			return err
		}
	})
}
