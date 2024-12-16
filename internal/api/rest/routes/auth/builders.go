package auth

import (
	"server/internal/entities"
	"server/pkg/app_error"
)

func (c *Controller) buildAuthMethodPlainOrThrow(dto BaseAuthDTO) (*entities.AuthMethodPlain, error) {
	authMethodPlain := c.buildAuthMethodPlain(dto)

	if err := authMethodPlain.Type.IsValid(); err != nil {
		return nil, app_error.New(err)
	}

	return &authMethodPlain, nil
}

func (c *Controller) buildAuthMethodPlain(dto BaseAuthDTO) entities.AuthMethodPlain {
	authMethodPlain := entities.AuthMethodPlain{}

	if dto.Email != "" {
		authMethodPlain.Type = entities.AuthMethodEmail
		authMethodPlain.Value = dto.Email
	}

	if dto.Phone != "" {
		authMethodPlain.Type = entities.AuthMethodPhone
		authMethodPlain.Value = dto.Phone
	}

	return authMethodPlain
}
