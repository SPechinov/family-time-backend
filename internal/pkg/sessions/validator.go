package sessions

import (
	"github.com/google/uuid"
	"server/internal/config"
	"server/internal/entities"
	"server/pkg/custom_error"
)

type Validator struct {
	cfg       *config.Config
	jwtParser *JWTParser
}

func NewValidator(cfg *config.Config) *Validator {
	return &Validator{
		cfg:       cfg,
		jwtParser: NewJWTParser(cfg),
	}
}

func (v *Validator) Validate(entity entities.SessionsValidate) error {
	token, err := v.jwtParser.Parse(entity.RefreshJWT)
	if err != nil || !token.Valid {
		return custom_error.ErrNotAuthorized
	}

	if err = v.validateSessionID(entity.SessionID); err != nil {
		return custom_error.ErrNotAuthorized
	}

	if entity.AccessJWT == nil {
		return nil
	}

	if err = v.validateJWT(*entity.AccessJWT); err != nil {
		return err
	}

	return nil
}

func (v *Validator) validateSessionID(sessionID string) error {
	if err := uuid.Validate(sessionID); err != nil {
		return custom_error.ErrNotAuthorized
	}

	return nil
}

func (v *Validator) validateJWT(jwt string) error {
	if token, err := v.jwtParser.Parse(jwt); err != nil || !token.Valid {
		return custom_error.ErrNotAuthorized
	}

	return nil
}
