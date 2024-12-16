package base_controller

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
	"server/internal/api/rest/utils/rest_error"
	"server/pkg/logger"
)

type Props[D any] struct {
	EchoCTX echo.Context
	Logger  *logger.Logger
	DTO     *D
}

type Handler[D any] func(props Props[D]) error

type BaseController[D any] struct {
	handler   Handler[D]
	validator func(*D) error
}

func New[D any](
	handler Handler[D],
	validator func(*D) error,
) *BaseController[D] {
	return &BaseController[D]{
		handler:   handler,
		validator: validator,
	}
}

func (baseController *BaseController[D]) Register(echoCTX echo.Context) error {
	l, err := baseController.getLogger(echoCTX)
	if err != nil {
		return err
	}

	if baseController.validator != nil {
		return baseController.withValidation(echoCTX, l)
	}

	return baseController.withoutValidation(echoCTX, l)
}

func (baseController *BaseController[D]) withoutValidation(echoCTX echo.Context, l *logger.Logger) error {
	return baseController.handler(Props[D]{
		EchoCTX: echoCTX,
		Logger:  l,
	})
}

func (baseController *BaseController[D]) withValidation(echoCTX echo.Context, l *logger.Logger) error {
	dto, err := baseController.getDTO(echoCTX, l)
	if err != nil {
		return err
	}

	err = baseController.validateDTO(dto, l)
	if err != nil {
		return err
	}

	return baseController.handler(Props[D]{
		EchoCTX: echoCTX,
		Logger:  l,
		DTO:     dto,
	})
}

func (baseController *BaseController[D]) getLogger(echoCTX echo.Context) (*logger.Logger, error) {
	l, err := echo_data.NewGetter(echoCTX).Logger()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (baseController *BaseController[D]) getDTO(echoCTX echo.Context, l *logger.Logger) (*D, error) {
	dto := new(D)
	err := echoCTX.Bind(dto)
	if err != nil {
		l.WithError(err).Debug("failed to bind request body")
		return nil, rest_error.ErrSomethingHappened
	}

	return dto, nil
}

func (baseController *BaseController[D]) validateDTO(dto *D, l *logger.Logger) error {
	err := baseController.validator(dto)
	if err != nil {
		l.WithError(err).Debug("failed validation")
		return err
	}

	return nil
}
