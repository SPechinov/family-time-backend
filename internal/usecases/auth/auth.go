package auth

import (
	"server/internal/config"
	"server/internal/entities"
	"server/pkg/custom_error"
	"server/pkg/logger"
	"sync"
)

type Auth struct {
	cfg             *config.Config
	usersService    usersService
	codesService    codesService
	messageSender   messageSender
	sessionsService sessionsService
	regConfirmMutex sync.Mutex
}

func New(
	cfg *config.Config,
	usersService usersService,
	codesService codesService,
	messageSender messageSender,
	sessionsService sessionsService,
) *Auth {
	return &Auth{
		cfg:             cfg,
		codesService:    codesService,
		usersService:    usersService,
		regConfirmMutex: sync.Mutex{},
		messageSender:   messageSender,
		sessionsService: sessionsService,
	}
}

func (a *Auth) Login(l *logger.Logger, entity entities.Login) (sessionData *entities.SessionData, err error) {
	l.WithAuthType(string(entity.AuthMethodPlain.Type))
	l.WithAuthValue(entity.AuthMethodPlain.Value)

	user, err := a.getUserByAuthMethodPlain(entity.AuthMethodPlain)
	if err != nil {
		return nil, err
	}

	passwordIsEqual := a.usersService.CompareHashAndPassword(user.Password, []byte(entity.Password))
	if !passwordIsEqual {
		return nil, custom_error.ErrIncorrectPassword
	}

	sessionData, err = a.sessionsService.Create(entities.SessionsCreate{UserID: user.UserID})
	if err != nil {
		return nil, err
	}

	return sessionData, nil
}

func (a *Auth) Registration(l *logger.Logger, entity entities.Registration) error {
	l.WithAuthType(string(entity.AuthMethodPlain.Type))
	l.WithAuthValue(entity.AuthMethodPlain.Value)

	err := a.ensureUserNotExistByAuthMethodPlain(entity.AuthMethodPlain)
	if err != nil {
		return err
	}

	code, err := a.generateAndSaveRegCode(entity.AuthMethodPlain.Value)
	if err != nil {
		return err
	}

	err = a.sendRegCodeToUser(entity.AuthMethodPlain, code)
	if err != nil {
		return err
	}

	l.WithConfirmationCode(code)
	return nil
}

func (a *Auth) RegistrationConfirm(l *logger.Logger, entity entities.RegistrationConfirm) error {
	a.regConfirmMutex.Lock()
	defer a.regConfirmMutex.Unlock()

	l.WithAuthType(string(entity.AuthMethodPlain.Type))
	l.WithAuthValue(entity.AuthMethodPlain.Value)
	l.WithConfirmationCode(entity.Code)

	err := a.ensureUserNotExistByAuthMethodPlain(entity.AuthMethodPlain)
	if err != nil {
		return err
	}

	err = a.checkRegCode(entity.AuthMethodPlain.Value, entity.Code)
	if err != nil {
		return err
	}

	userCreateSpec := a.buildUserCreateSpec(entity)

	_, err = a.usersService.Create(userCreateSpec)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth) ForgotPassword(l *logger.Logger, entity entities.ForgotPassword) error {
	l.WithAuthType(string(entity.AuthMethodPlain.Type))
	l.WithAuthValue(entity.AuthMethodPlain.Value)

	err := a.ensureUserExistByAuthMethodPlain(entity.AuthMethodPlain)
	if err != nil {
		return err
	}

	code, err := a.generateAndSaveForgotPasswordCode(entity.AuthMethodPlain.Value)
	if err != nil {
		return err
	}

	err = a.sendForgotPasswordCodeToUser(entity.AuthMethodPlain, code)
	if err != nil {
		return err
	}

	l.WithConfirmationCode(code)
	return nil
}

func (a *Auth) ForgotPasswordConfirm(l *logger.Logger, entity entities.ForgotPasswordConfirm) error {
	l.WithAuthType(string(entity.AuthMethodPlain.Type))
	l.WithAuthValue(entity.AuthMethodPlain.Value)
	l.WithConfirmationCode(entity.Code)

	user, err := a.getUserByAuthMethodPlain(entity.AuthMethodPlain)
	if err != nil {
		return err
	}

	err = a.checkForgotPasswordCode(entity.AuthMethodPlain.Value, entity.Code)
	if err != nil {
		return err
	}

	userPatchOneSpec := a.buildForgotPasswordSpec(entity)
	_, err = a.usersService.PatchOne(userPatchOneSpec)
	if err != nil {
		return err
	}

	err = a.sessionsService.DeleteAll(user.UserID)
	if err != nil {
		l.Error(err)
	}

	return nil
}

func (a *Auth) UpdateSession(entity entities.UpdateSession) (sessionData *entities.SessionData, err error) {
	err = a.ValidateSession(entities.SessionsValidate{
		SessionID:  entity.SessionID,
		RefreshJWT: entity.RefreshJWT,
	})
	if err != nil {
		return nil, err
	}

	userID, err := a.GetUserIdFromJWT(entity.RefreshJWT)
	if err != nil {
		return nil, err
	}

	err = a.sessionsService.HasSessionInStore(entities.SessionsHas{
		UserID:     userID,
		SessionID:  entity.SessionID,
		RefreshJWT: entity.RefreshJWT,
	})
	if err != nil {
		return nil, err
	}

	err = a.ensureUserExistByUserID(userID)
	if err != nil {
		return nil, err
	}

	err = a.sessionsService.Delete(entities.Logout{
		UserID:    userID,
		SessionID: entity.SessionID,
	})
	if err != nil {
		return nil, err
	}

	newSessionData, err := a.sessionsService.Create(entities.SessionsCreate{UserID: userID})
	if err != nil {
		return nil, err
	}

	return newSessionData, nil
}

func (a *Auth) Logout(entity entities.Logout) error {
	return a.sessionsService.Delete(entity)
}
func (a *Auth) LogoutAll(userID string) error {
	return a.sessionsService.DeleteAll(userID)
}
