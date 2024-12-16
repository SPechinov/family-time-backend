package middlewares

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/utils/echo_data"
	"server/internal/api/rest/utils/sessions"
	"server/internal/config"
	"server/internal/entities"
	pkgSessions "server/internal/pkg/sessions"
)

func CheckAuth(cfg *config.Config) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCTX echo.Context) error {
			sessionInfo, err := getSessionInfoFromEcho(echoCTX)
			if err != nil {
				return err
			}

			err = pkgSessions.NewValidator(cfg).Validate(*sessionInfo)
			if err != nil {
				return err
			}

			userID, err := getUserIDFromJWT(cfg, sessionInfo.RefreshJWT)
			if err != nil {
				return err
			}

			echoSetter := echo_data.NewSetter(echoCTX)
			echoSetter.UserID(userID)
			echoSetter.SessionID(sessionInfo.SessionID)

			err = logUserID(echoCTX, userID)
			if err != nil {
				return err
			}

			return next(echoCTX)
		}
	}
}

func getSessionInfoFromEcho(echoCTX echo.Context) (entity *entities.SessionsValidate, err error) {
	sessionGetter := sessions.NewSessionGetter(echoCTX)

	sessionID, err := sessionGetter.GetSessionID()
	if err != nil {
		return nil, err
	}

	refreshJWT, err := sessionGetter.GetRefreshJWT()
	if err != nil {
		return nil, err
	}

	accessJWT, err := sessionGetter.GetAccessJWT()
	if err != nil {
		return nil, err
	}

	return &entities.SessionsValidate{
		SessionID:  sessionID,
		RefreshJWT: refreshJWT,
		AccessJWT:  &accessJWT,
	}, nil
}

func getUserIDFromJWT(cfg *config.Config, jwt string) (userID string, err error) {
	jwtParser := pkgSessions.NewJWTParser(cfg)
	if _, err = jwtParser.Parse(jwt); err != nil {
		return "", err
	}

	userID, err = jwtParser.GetUserID()
	if err != nil {
		return "", err
	}

	return userID, nil
}

func logUserID(echoCTX echo.Context, userID string) error {
	l, err := echo_data.NewGetter(echoCTX).Logger()
	if err != nil {
		return err
	}

	l.WithUserID(userID)
	return nil
}
