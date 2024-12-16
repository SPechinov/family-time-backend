package auth

import (
	"server/internal/entities"
	"server/pkg/logger"
)

type useCases interface {
	Login(l *logger.Logger, entity entities.Login) (sessionData *entities.SessionData, err error)
	Registration(l *logger.Logger, entity entities.Registration) (err error)
	RegistrationConfirm(l *logger.Logger, entity entities.RegistrationConfirm) error
	ForgotPassword(l *logger.Logger, entity entities.ForgotPassword) error
	ForgotPasswordConfirm(l *logger.Logger, entity entities.ForgotPasswordConfirm) error
	UpdateSession(entity entities.UpdateSession) (sessionData *entities.SessionData, err error)
	Logout(entity entities.Logout) error
	LogoutAll(userID string) error
}
