package constants

import "time"

const (
	SessionAccessJWTDuration  = time.Minute * 5
	SessionRefreshJWTDuration = time.Hour * 720
)
