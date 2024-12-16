package constants

import "time"

const (
	RegistrationCodeLiveTime    = time.Minute * 10
	RegistrationCodeMaxAttempts = 5

	ForgotPasswordCodeLiveTime    = time.Minute * 10
	ForgotPasswordCodeMaxAttempts = 5
)
