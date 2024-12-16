package entities

type SessionData struct {
	UserID     string
	SessionID  string
	RefreshJWT string
	AccessJWT  string
}

type SessionsCreate struct {
	UserID string
}

type SessionsHas struct {
	UserID     string
	SessionID  string
	RefreshJWT string
}

type SessionsValidate struct {
	SessionID  string
	RefreshJWT string
	AccessJWT  *string
}
