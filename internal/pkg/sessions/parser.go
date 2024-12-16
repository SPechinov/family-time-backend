package sessions

import (
	"github.com/golang-jwt/jwt/v5"
	"server/internal/config"
	"server/pkg/custom_error"
)

type JWTParser struct {
	cfg   *config.Config
	token *jwt.Token
}

func NewJWTParser(cfg *config.Config) *JWTParser {
	return &JWTParser{
		cfg: cfg,
	}
}

func (p *JWTParser) Parse(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(p.cfg.Auth.JWTSecretKey), nil
	})
	if err != nil {
		return nil, custom_error.ErrNotAuthorized
	}

	p.token = jwtToken
	return jwtToken, nil
}

func (p *JWTParser) GetUserID() (string, error) {
	userID, ok := p.token.Claims.(jwt.MapClaims)[JWTKeyUserID].(string)
	if !ok {
		return "", custom_error.ErrNotAuthorized
	}
	return userID, nil
}
