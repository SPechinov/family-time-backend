package sessions

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"server/internal/constants"
	"server/internal/entities"
	"server/pkg/app_error"
)

type Sessions struct {
	redis *redis.Client
}

func New(redis *redis.Client) *Sessions {
	return &Sessions{
		redis: redis,
	}
}

func (s *Sessions) Add(entity entities.SessionData) error {
	ketSession := getKeySession(entity.UserID, entity.SessionID)
	err := s.redis.Set(context.TODO(), ketSession, entity.RefreshJWT, constants.SessionRefreshJWTDuration).Err()
	if err != nil {
		return app_error.New(err)
	}
	return nil
}

func (s *Sessions) Get(userID, sessionID string) (string, error) {
	value, err := s.redis.Get(context.TODO(), getKeySession(userID, sessionID)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return "", app_error.New(err)
	}
	return value, nil
}

func (s *Sessions) Delete(entity entities.Logout) error {
	err := s.redis.Del(context.TODO(), getKeySession(entity.UserID, entity.SessionID)).Err()
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (s *Sessions) DeleteAll(userID string) error {
	var cursor uint64
	var err error

	for {
		keys, newCursor, scanErr := s.redis.Scan(context.TODO(), cursor, getKeyAllUserSession(userID), 0).Result()
		err = scanErr

		for _, key := range keys {
			delErr := s.redis.Del(context.TODO(), key).Err()
			err = delErr
		}

		if newCursor == 0 {
			break
		}

		cursor = newCursor
	}

	return err
}

func getKeySession(userID, sessionID string) string {
	return "sessions:" + userID + ":" + sessionID
}

func getKeyAllUserSession(userID string) string {
	return "sessions:" + userID + ":*"
}
