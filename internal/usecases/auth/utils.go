package auth

import (
	"errors"
	"server/internal/constants"
	"server/internal/constants/validation_rules"
	"server/internal/entities"
	"server/internal/pkg/sessions"
	"server/pkg/app_error"
	"server/pkg/custom_error"
	"server/pkg/utils"
)

func (a *Auth) getUserByAuthMethodPlain(authMethodPlain entities.AuthMethodPlain) (*entities.User, error) {
	spec := a.buildUserSearchSpec(authMethodPlain)
	user, err := a.usersService.FindOne(spec)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, custom_error.ErrUserNotExists
	}

	return user, nil
}

func (a *Auth) ensureUserNotExistByAuthMethodPlain(authMethodPlain entities.AuthMethodPlain) error {
	spec := a.buildUserSearchSpec(authMethodPlain)

	exists, err := a.usersService.Exists(spec)
	if err != nil {
		return err
	}

	if exists {
		return custom_error.ErrUserExists
	}

	return nil
}

func (a *Auth) ensureUserExistByAuthMethodPlain(authMethodPlain entities.AuthMethodPlain) error {
	spec := a.buildUserSearchSpec(authMethodPlain)

	exists, err := a.usersService.Exists(spec)
	if err != nil {
		return err
	}

	if !exists {
		return custom_error.ErrUserNotExists
	}

	return nil
}

func (a *Auth) ensureUserExistByUserID(userID string) error {
	entity := entities.UsersFindOneSpec{UserID: &userID, Deleted: false}

	exists, err := a.usersService.Exists(entity)
	if err != nil {
		return err
	}

	if !exists {
		return custom_error.ErrUserNotExists
	}

	return nil
}

func (a *Auth) generateAndSaveRegCode(authMethod string) (string, error) {
	code := utils.GenerateRandomCode(validation_rules.LenRegistrationCode)

	err := a.codesService.Save(
		a.generateRedisKeyForRegCode(authMethod),
		code,
		constants.RegistrationCodeLiveTime,
	)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (a *Auth) checkRegCode(authMethodValue string, code string) error {
	return a.codesService.CompareCodesAndIncrementOrDeleteIfNotEqual(
		a.generateRedisKeyForRegCode(authMethodValue),
		code,
		constants.RegistrationCodeMaxAttempts,
	)
}

func (a *Auth) sendRegCodeToUser(authMethod entities.AuthMethodPlain, code string) error {
	sendMethod, err := a.mapAuthMethodTypeToSendMethod(authMethod.Type)
	if err != nil {
		return err
	}

	err = a.messageSender.SendRegMessage(sendMethod, authMethod.Value, code)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth) generateRedisKeyForRegCode(key string) string {
	return "rest:auth:reg-code:" + key
}

func (a *Auth) generateAndSaveForgotPasswordCode(authMethod string) (string, error) {
	code := utils.GenerateRandomCode(validation_rules.LenForgotPasswordCode)

	err := a.codesService.Save(
		a.generateRedisKeyForForgotPasswordCode(authMethod),
		code,
		constants.ForgotPasswordCodeLiveTime,
	)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (a *Auth) checkForgotPasswordCode(authMethodValue string, code string) error {
	return a.codesService.CompareCodesAndIncrementOrDeleteIfNotEqual(
		a.generateRedisKeyForForgotPasswordCode(authMethodValue),
		code,
		constants.ForgotPasswordCodeMaxAttempts,
	)
}

func (a *Auth) sendForgotPasswordCodeToUser(authMethod entities.AuthMethodPlain, code string) error {
	sendMethod, err := a.mapAuthMethodTypeToSendMethod(authMethod.Type)
	if err != nil {
		return err
	}

	err = a.messageSender.SendForgotPasswordMessage(sendMethod, authMethod.Value, code)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth) generateRedisKeyForForgotPasswordCode(key string) string {
	return "rest:auth:forgot-password-code:" + key
}

func (a *Auth) mapAuthMethodTypeToSendMethod(authMethodtype entities.AuthMethodType) (entities.MessageMethod, error) {
	var method entities.MessageMethod

	if err := authMethodtype.IsValid(); err != nil {
		return method, app_error.New(err)
	}

	switch authMethodtype {
	case entities.AuthMethodEmail:
		method = entities.MessageMethodEmail
	case entities.AuthMethodPhone:
		method = entities.MessageMethodPhone
	default:
		return method, app_error.New(errors.New("unknown auth method type"))
	}

	return method, nil
}

func (a *Auth) ValidateSession(entity entities.SessionsValidate) error {
	jwtParser := sessions.NewJWTParser(a.cfg)
	token, err := jwtParser.Parse(entity.RefreshJWT)
	if err != nil || !token.Valid {
		return custom_error.ErrNotAuthorized
	}

	sessionValidator := sessions.NewValidator(a.cfg)
	return sessionValidator.Validate(entity)
}

func (a *Auth) GetUserIdFromJWT(jwtToken string) (userID string, err error) {
	jwtParser := sessions.NewJWTParser(a.cfg)
	_, err = jwtParser.Parse(jwtToken)
	if err != nil {
		return "", err
	}

	return jwtParser.GetUserID()
}
