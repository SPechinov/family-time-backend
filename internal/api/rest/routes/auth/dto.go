package auth

type BaseAuthDTO struct {
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

type LoginDTO struct {
	BaseAuthDTO
	Password string `json:"password" form:"password"`
}

type RegistrationDTO struct {
	BaseAuthDTO
}

type RegistrationConfirmDTO struct {
	BaseAuthDTO
	FirstName   string `json:"firstName" form:"firstName"`
	Password    string `json:"password" form:"password"`
	CountryCode string `json:"countryCode" form:"countryCode"`
	Code        string `json:"code" form:"code"`
}

type ForgotPasswordDTO struct {
	BaseAuthDTO
}

type ForgotPasswordConfirmDTO struct {
	BaseAuthDTO
	Password string `json:"password" form:"password"`
	Code     string `json:"code" form:"code"`
}
