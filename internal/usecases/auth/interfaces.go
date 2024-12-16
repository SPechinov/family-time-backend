package auth

import (
	"server/internal/entities"
	"time"
)

type codesService interface {
	Save(key string, code string, ttl time.Duration) error
	CompareCodesAndIncrementOrDeleteIfNotEqual(key string, userCode string, maxAttempts int) error
}

type usersService interface {
	Create(spec entities.UserCreateSpec) (user *entities.User, err error)
	Exists(spec entities.UsersFindOneSpec) (bool, error)
	FindOne(spec entities.UsersFindOneSpec) (user *entities.User, err error)
	FindMany(spec entities.UsersFindManySpec) (users []entities.User, err error)
	PatchOne(spec entities.UsersPatchOneSpec) (*entities.User, error)
	CompareHashAndPassword(hash, password []byte) bool
}

type messageSender interface {
	SendRegMessage(messageMethod entities.MessageMethod, recipient string, code string) error
	SendForgotPasswordMessage(messageMethod entities.MessageMethod, recipient string, code string) error
}

type sessionsService interface {
	Create(entity entities.SessionsCreate) (sessionData *entities.SessionData, err error)
	HasSessionInStore(entity entities.SessionsHas) error
	Delete(entity entities.Logout) error
	DeleteAll(userID string) error
}
