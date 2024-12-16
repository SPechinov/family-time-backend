package sessions

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SessionCleaner struct {
	echoCTX echo.Context
}

func NewSessionCleaner(echoCTX echo.Context) *SessionCleaner {
	return &SessionCleaner{
		echoCTX: echoCTX,
	}
}

func (c *SessionCleaner) Clean() {
	c.cleanRefreshJWT()
	c.cleanSessionID()
}

func (c *SessionCleaner) cleanRefreshJWT() {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieRefreshJWT,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}

func (c *SessionCleaner) cleanSessionID() {
	c.echoCTX.SetCookie(&http.Cookie{
		Name:     cookieSessionID,
		Value:    "",
		Path:     "/",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
}
