package echo_data

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"server/internal/api/rest/utils/rest_error"
	"server/pkg/logger"
)

type Getter struct {
	echoCTX echo.Context
}

func NewGetter(echoCTX echo.Context) *Getter {
	return &Getter{
		echoCTX: echoCTX,
	}
}

func (g *Getter) Logger() (l *logger.Logger, err error) {
	l, ok := g.echoCTX.Get(keyLogger).(*logger.Logger)
	if !ok {
		logrus.Error("Logger not in echo context")
		return nil, rest_error.ErrSomethingHappened
	}

	return
}

func (g *Getter) UserID() (userID string, err error) {
	userID, ok := g.echoCTX.Get(keyUserID).(string)
	if !ok {
		logrus.Error("UserID not in echo context")
		return "", rest_error.ErrSomethingHappened
	}
	return
}

func (g *Getter) SessionID() (sessionID string, err error) {
	sessionID, ok := g.echoCTX.Get(keySessionID).(string)
	if !ok {
		logrus.Error("SessionID not in echo context")
		return "", rest_error.ErrSomethingHappened
	}
	return
}

func (g *Getter) RequestID() (requestID string, err error) {
	requestID, ok := g.echoCTX.Get(keyRequestID).(string)
	if !ok {
		logrus.Error("RequestID not in echo context")
		return "", rest_error.ErrSomethingHappened
	}
	return
}
