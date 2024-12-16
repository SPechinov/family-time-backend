package auth

import (
	"server/internal/entities"
)

func (a *Auth) buildUserSearchSpec(
	authMethodPlain entities.AuthMethodPlain,
) entities.UsersFindOneSpec {
	spec := entities.UsersFindOneSpec{}

	value := []byte(authMethodPlain.Value)

	switch authMethodPlain.Type {
	case entities.AuthMethodEmail:
		spec.Email = &value
	case entities.AuthMethodPhone:
		spec.Phone = &value
	}

	return spec
}

func (a *Auth) buildUserCreateSpec(entity entities.RegistrationConfirm) entities.UserCreateSpec {
	authMethod := a.buildAuthMethodSpec(entity.AuthMethodPlain)

	userCreateSpec := entities.UserCreateSpec{
		AuthMethodSpec: authMethod,
		Password:       []byte(entity.Password),
		FirstName:      entity.FirstName,
		CountryCode:    entity.CountryCode,
	}

	return userCreateSpec
}

func (a *Auth) buildAuthMethodSpec(entity entities.AuthMethodPlain) entities.AuthMethodSpec {
	authMethod := entities.AuthMethodSpec{}

	switch entity.Type {
	case entities.AuthMethodEmail:
		authMethod.Type = entities.AuthMethodEmail
	case entities.AuthMethodPhone:
		authMethod.Type = entities.AuthMethodPhone
	}

	authMethod.Values.Searchable = []byte(entity.Value)
	authMethod.Values.Encrypted = []byte(entity.Value)

	return authMethod
}

func (a *Auth) buildForgotPasswordSpec(entity entities.ForgotPasswordConfirm) entities.UsersPatchOneSpec {
	spec := entities.UsersPatchOneSpec{
		Data: entities.UsersPatchDataSpec{},
	}

	value := []byte(entity.AuthMethodPlain.Value)
	switch entity.AuthMethodPlain.Type {
	case entities.AuthMethodEmail:
		spec.Email = &value
	case entities.AuthMethodPhone:
		spec.Phone = &value
	}

	password := []byte(entity.Password)
	spec.Data.Password = &password

	return spec
}
