package auth

import (
	"net/http"
	"server/internal/api/rest/types"
	"server/internal/api/rest/utils/base_controller"
	"server/internal/api/rest/utils/echo_data"
	"server/internal/api/rest/utils/sessions"
	"server/internal/entities"
)

func (c *Controller) login(props base_controller.Props[LoginDTO]) error {
	authMethodPlain, err := c.buildAuthMethodPlainOrThrow(props.DTO.BaseAuthDTO)
	if err != nil {
		return err
	}

	entity := entities.Login{
		AuthMethodPlain: *authMethodPlain,
		Password:        props.DTO.Password,
	}

	sessionData, err := c.useCases.Login(props.Logger, entity)
	if err != nil {
		return err
	}

	sessionSetter := sessions.NewSessionSetter(props.EchoCTX)
	sessionSetter.SetSessionID(sessionData.SessionID)
	sessionSetter.SetRefreshJWT(sessionData.RefreshJWT)
	sessionSetter.SetAccessJWT(sessionData.AccessJWT)

	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) registration(props base_controller.Props[RegistrationDTO]) error {
	authMethodPlain, err := c.buildAuthMethodPlainOrThrow(props.DTO.BaseAuthDTO)
	if err != nil {
		return err
	}

	entity := entities.Registration{
		AuthMethodPlain: *authMethodPlain,
	}

	err = c.useCases.Registration(props.Logger, entity)
	if err != nil {
		return err
	}

	return props.EchoCTX.NoContent(http.StatusOK)
}

func (c *Controller) registrationConfirm(props base_controller.Props[RegistrationConfirmDTO]) error {
	authMethodPlain, err := c.buildAuthMethodPlainOrThrow(props.DTO.BaseAuthDTO)
	if err != nil {
		return err
	}

	entity := entities.RegistrationConfirm{
		AuthMethodPlain: *authMethodPlain,
		Password:        props.DTO.Password,
		CountryCode:     props.DTO.CountryCode,
		Code:            props.DTO.Code,
	}

	err = c.useCases.RegistrationConfirm(props.Logger, entity)
	if err != nil {
		return err
	}

	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) forgotPassword(props base_controller.Props[ForgotPasswordDTO]) error {
	authMethodPlain, err := c.buildAuthMethodPlainOrThrow(props.DTO.BaseAuthDTO)
	if err != nil {
		return err
	}

	entity := entities.ForgotPassword{
		AuthMethodPlain: *authMethodPlain,
	}

	err = c.useCases.ForgotPassword(props.Logger, entity)
	if err != nil {
		return err
	}

	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) forgotPasswordConfirm(props base_controller.Props[ForgotPasswordConfirmDTO]) error {
	authMethodPlain, err := c.buildAuthMethodPlainOrThrow(props.DTO.BaseAuthDTO)
	if err != nil {
		return err
	}

	entity := entities.ForgotPasswordConfirm{
		AuthMethodPlain: *authMethodPlain,
		Code:            props.DTO.Code,
		Password:        props.DTO.Password,
	}

	err = c.useCases.ForgotPasswordConfirm(props.Logger, entity)
	if err != nil {
		return err
	}

	sessions.NewSessionCleaner(props.EchoCTX).Clean()
	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) updateJWT(props base_controller.Props[types.EmptyDTO]) error {
	sessionGetter := sessions.NewSessionGetter(props.EchoCTX)

	sessionID, err := sessionGetter.GetSessionID()
	if err != nil {
		return err
	}
	props.Logger.WithField("session_id", sessionID)

	refreshJWT, err := sessionGetter.GetRefreshJWT()
	if err != nil {
		return err
	}
	props.Logger.WithField("refresh_jwt", refreshJWT)

	sessionData, err := c.useCases.UpdateSession(entities.UpdateSession{
		SessionID:  sessionID,
		RefreshJWT: refreshJWT,
	})
	if err != nil {
		return err
	}

	sessionSetter := sessions.NewSessionSetter(props.EchoCTX)
	sessionSetter.SetSessionID(sessionData.SessionID)
	sessionSetter.SetRefreshJWT(sessionData.RefreshJWT)
	sessionSetter.SetAccessJWT(sessionData.AccessJWT)

	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) logout(props base_controller.Props[types.EmptyDTO]) error {
	echoGetter := echo_data.NewGetter(props.EchoCTX)
	userID, err := echoGetter.UserID()
	if err != nil {
		return err
	}

	sessionID, err := echoGetter.SessionID()
	if err != nil {
		return err
	}

	err = c.useCases.Logout(entities.Logout{
		UserID:    userID,
		SessionID: sessionID,
	})
	if err != nil {
		return err
	}

	sessionSetter := sessions.NewSessionSetter(props.EchoCTX)
	sessionSetter.ClearSessionID()
	sessionSetter.ClearRefreshJWT()

	return props.EchoCTX.NoContent(http.StatusNoContent)
}

func (c *Controller) logoutAll(props base_controller.Props[types.EmptyDTO]) error {
	echoGetter := echo_data.NewGetter(props.EchoCTX)
	userID, err := echoGetter.UserID()
	if err != nil {
		return err
	}

	err = c.useCases.LogoutAll(userID)
	if err != nil {
		return err
	}

	sessionSetter := sessions.NewSessionSetter(props.EchoCTX)
	sessionSetter.ClearSessionID()
	sessionSetter.ClearRefreshJWT()

	return props.EchoCTX.NoContent(http.StatusNoContent)
}
