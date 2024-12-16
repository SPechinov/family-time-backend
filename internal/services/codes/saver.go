package codes

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/pkg/app_error"
	"time"
)

type saver struct {
	redis *redis.Client
}

func newSaver(redis *redis.Client) *saver {
	return &saver{
		redis: redis,
	}
}

func (s *saver) Save(key string, code string, attempts int, ttl time.Duration) error {
	redisValue := s.buildRedisValue(code, attempts)
	err := s.redis.Set(context.TODO(), key, redisValue, ttl).Err()
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (s *saver) buildRedisValue(code string, attempts int) string {
	return fmt.Sprintf(codeFormat, code, attempts)
}
