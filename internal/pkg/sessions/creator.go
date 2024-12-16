package sessions

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"server/internal/config"
	"server/internal/constants"
	"server/internal/entities"
	"server/pkg/app_error"
	"time"
)

type Creator struct {
	cfg *config.Config
}

func NewCreator(cfg *config.Config) *Creator {
	return &Creator{
		cfg: cfg,
	}
}

func (c *Creator) Create(entity entities.SessionsCreate) (*entities.SessionData, error) {
	refreshJWT, accessJWT, err := c.createJWTPair(entity.UserID)
	if err != nil {
		return nil, err
	}
	sessionID := c.createSessionID()

	return &entities.SessionData{
		UserID:     entity.UserID,
		SessionID:  sessionID,
		RefreshJWT: refreshJWT,
		AccessJWT:  accessJWT,
	}, nil
}

func (c *Creator) createJWTPair(userID string) (refreshJWT, accessJWT string, err error) {
	refreshJWT, err = c.createRefreshJWT(userID)
	if err != nil {
		return "", "", err
	}

	accessJWT, err = c.createAccessJWT(userID)
	if err != nil {
		return "", "", err
	}

	return refreshJWT, accessJWT, nil
}

func (c *Creator) createAccessJWT(userID string) (string, error) {
	payload := c.createJWTPayload(userID, JWTTypeAccess)
	return c.createJWT(payload)
}

func (c *Creator) createRefreshJWT(userID string) (string, error) {
	payload := c.createJWTPayload(userID, JWTTypeRefresh)
	return c.createJWT(payload)
}

func (c *Creator) createJWTPayload(userID string, jwtType string) jwt.MapClaims {
	duration := c.createJWTDuration(jwtType)

	mapClaims := jwt.MapClaims{
		"exp":        time.Now().Add(duration).Unix(),
		JWTKeyUserID: userID,
	}

	return mapClaims
}

func (c *Creator) createJWTDuration(jwtType string) time.Duration {
	if jwtType == JWTTypeAccess {
		return constants.SessionAccessJWTDuration
	}
	return constants.SessionRefreshJWTDuration
}

func (c *Creator) createJWT(mapClaims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, mapClaims)

	jwtString, err := token.SignedString([]byte(c.cfg.Auth.JWTSecretKey))
	if err != nil {
		return "", app_error.New(err)
	}

	return jwtString, nil
}

func (c *Creator) createSessionID() string {
	return uuid.New().String()
}
