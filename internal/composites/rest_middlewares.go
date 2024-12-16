package composites

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/middlewares"
)

func NewRestMiddlewares(restGroup *echo.Group) {
	config := middlewares.Config{
		EnableRequestID: true,
		EnableLogURL:    true,
	}
	middlewares.New(&config).Register(restGroup)
}
