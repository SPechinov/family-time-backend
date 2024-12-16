package sessions

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/rest_error"
	"strings"
)

type SessionGetter struct {
	echoCTX echo.Context
}

func NewSessionGetter(echoCTX echo.Context) *SessionGetter {
	return &SessionGetter{
		echoCTX: echoCTX,
	}
}

func (c *SessionGetter) GetSessionID() (string, error) {
	refreshJWTCookie, err := c.echoCTX.Cookie(cookieSessionID)
	if err != nil || len(refreshJWTCookie.Value) == 0 {
		return "", rest_error.ErrNotAuthorized
	}

	return refreshJWTCookie.Value, nil
}

func (c *SessionGetter) GetAccessJWT() (string, error) {
	accessJWTDirty := c.echoCTX.Request().Header[headerAccessJWT]

	if len(accessJWTDirty) == 0 || len(accessJWTDirty[0]) == 0 {
		return "", rest_error.ErrNotAuthorized
	}

	return strings.TrimPrefix(accessJWTDirty[0], "Bearer "), nil
}

func (c *SessionGetter) GetRefreshJWT() (string, error) {
	refreshJWTCookie, err := c.echoCTX.Cookie(cookieRefreshJWT)
	if err != nil || len(refreshJWTCookie.Value) == 0 {
		return "", rest_error.ErrNotAuthorized
	}

	return refreshJWTCookie.Value, nil
}
