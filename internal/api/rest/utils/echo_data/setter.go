package echo_data

import (
	"github.com/labstack/echo/v4"
	"server/pkg/logger"
)

type Setter struct {
	echoCTX echo.Context
}

func NewSetter(echoCTX echo.Context) *Setter {
	return &Setter{
		echoCTX: echoCTX,
	}
}

func (s *Setter) Logger(l *logger.Logger) {
	s.echoCTX.Set(keyLogger, l)
}

func (s *Setter) UserID(userID string) {
	s.echoCTX.Set(keyUserID, userID)
}

func (s *Setter) SessionID(sessionID string) {
	s.echoCTX.Set(keySessionID, sessionID)
}

func (s *Setter) RequestID(requestID string) {
	s.echoCTX.Response().Header().Set(keyHeaderXRequestID, requestID)
	s.echoCTX.Set(keyRequestID, requestID)
}
