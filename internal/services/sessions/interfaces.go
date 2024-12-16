package sessions

import "server/internal/entities"

type JWTType string

type sessionsStore interface {
	Add(entity entities.SessionData) error
	Get(userID, sessionID string) (string, error)
	Delete(entity entities.Logout) error
	DeleteAll(userID string) error
}
