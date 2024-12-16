package auth

import (
	"github.com/labstack/echo/v4"
	"server/internal/api/rest/middlewares"
	"server/internal/api/rest/types"
	"server/internal/api/rest/utils/base_controller"
	"server/internal/config"
)

const (
	urlLogin                 = "/login"
	urlRegistration          = "/registration"
	urlRegistrationConfirm   = "/registration-confirm"
	urlForgotPassword        = "/forgot-password"
	urlForgotPasswordConfirm = "/forgot-password-confirm"
	urlUpdateJWT             = "/update-jwt"
	urlLogout                = "/logout"
	urlLogoutAll             = "/logout-all"
)

type Controller struct {
	cfg      *config.Config
	useCases useCases
}

func New(cfg *config.Config, useCases useCases) *Controller {
	return &Controller{
		cfg:      cfg,
		useCases: useCases,
	}
}

func (c *Controller) Register(restGroup *echo.Group) {
	authRestGroup := restGroup.Group("/auth")

	authRestGroup.POST(
		urlLogin,
		base_controller.New(c.login, validateLoginDTO).Register,
	)
	authRestGroup.POST(
		urlRegistration,
		base_controller.New(c.registration, validateRegistrationDTO).Register,
	)
	authRestGroup.POST(
		urlRegistrationConfirm,
		base_controller.New(c.registrationConfirm, validateRegistrationConfirmDTO).Register,
	)
	authRestGroup.POST(
		urlForgotPassword,
		base_controller.New(c.forgotPassword, validateForgotPasswordDTO).Register,
	)
	authRestGroup.POST(
		urlForgotPasswordConfirm,
		base_controller.New(c.forgotPasswordConfirm, validateForgotPasswordConfirmDTO).Register,
	)

	authRestGroup.POST(
		urlUpdateJWT,
		base_controller.New[types.EmptyDTO](c.updateJWT, nil).Register,
	)

	authProtectedRestGroup := authRestGroup.Group("", middlewares.CheckAuth(c.cfg))

	authProtectedRestGroup.POST(
		urlLogout,
		base_controller.New[types.EmptyDTO](c.logout, nil).Register,
	)

	authProtectedRestGroup.POST(
		urlLogoutAll,
		base_controller.New[types.EmptyDTO](c.logoutAll, nil).Register,
	)
}
