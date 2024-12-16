package codes

import (
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	codeFormat = "%s : %d"
)

type Codes struct {
	redis    *redis.Client
	saver    *saver
	comparer *comparer
}

func New(redis *redis.Client) *Codes {
	cs := newSaver(redis)
	cc := newComparer(redis, cs)

	return &Codes{
		redis:    redis,
		saver:    cs,
		comparer: cc,
	}
}

func (c *Codes) Save(key string, code string, ttl time.Duration) error {
	return c.saver.Save(key, code, 0, ttl)
}

func (c *Codes) CompareCodesAndIncrementOrDeleteIfNotEqual(key string, userCode string, maxAttempts int) error {
	return c.comparer.CompareCodesAndIncrementOrDeleteIfNotEqual(key, userCode, maxAttempts)
}
