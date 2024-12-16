package middlewares

import (
	"github.com/labstack/echo/v4"
	mwResponseHandler "server/internal/api/rest/middlewares/response_handler"
)

type Config struct {
	EnableLogURL    bool
	EnableRequestID bool
}

type Middlewares struct {
	config *Config
}

func New(config *Config) *Middlewares {
	return &Middlewares{
		config: config,
	}
}

func (middlewares *Middlewares) Register(restGroup *echo.Group) {
	registerLogger(restGroup)

	if middlewares.config.EnableRequestID {
		registerRequestID(restGroup)
		registerLogRequestID(restGroup)
	}

	if middlewares.config.EnableLogURL {
		registerLogURI(restGroup)
	}

	mwResponseHandler.Register(restGroup)
}
