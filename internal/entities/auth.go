package entities

import (
	"errors"
	"fmt"
)

type AuthMethodType string

const (
	AuthMethodEmail AuthMethodType = "email"
	AuthMethodPhone AuthMethodType = "phone"
)

func (authMethodType *AuthMethodType) IsValid() error {
	switch *authMethodType {
	case AuthMethodEmail, AuthMethodPhone:
		return nil
	}

	return errors.New(fmt.Sprintf("%s - unknown auth method type", string(*authMethodType)))
}

type Login struct {
	AuthMethodPlain
	Password string
}

type Registration struct {
	AuthMethodPlain
}

type RegistrationConfirm struct {
	AuthMethodPlain
	FirstName   string
	Password    string
	CountryCode string
	Code        string
}

type ForgotPassword struct {
	AuthMethodPlain
}

type ForgotPasswordConfirm struct {
	AuthMethodPlain
	Code     string
	Password string
}

type UpdateSession struct {
	SessionID  string
	RefreshJWT string
}

type Logout struct {
	UserID    string
	SessionID string
}

type AuthMethodPlain struct {
	Type  AuthMethodType
	Value string
}

type AuthMethodValues struct {
	Encrypted  []byte
	Searchable []byte
}
