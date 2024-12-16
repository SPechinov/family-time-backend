package sessions

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"server/internal/constants"
	"time"
)

type SessionSetter struct {
	echoCTX echo.Context
}

func NewSessionSetter(echoCTX echo.Context) *SessionSetter {
	return &SessionSetter{
		echoCTX: echoCTX,
	}
}

func (c *SessionSetter) SetAccessJWT(accessJWT string) {
	c.echoCTX.Response().Header().Add(headerAccessJWT, accessJWT)
}

func (c *SessionSetter) SetRefreshJWT(refreshJWT string) {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieRefreshJWT,
		Value:    refreshJWT,
		Path:     "/",
		MaxAge:   int(constants.SessionRefreshJWTDuration / time.Second),
		Secure:   true,
		HttpOnly: true,
	})
}

func (c *SessionSetter) SetSessionID(sessionID string) {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieSessionID,
		Value:    sessionID,
		Path:     "/",
		MaxAge:   int(constants.SessionRefreshJWTDuration / time.Second),
		Secure:   true,
		HttpOnly: true,
	})
}

func (c *SessionSetter) ClearRefreshJWT() {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieRefreshJWT,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}

func (c *SessionSetter) ClearSessionID() {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieSessionID,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}
