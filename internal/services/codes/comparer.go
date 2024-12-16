package codes

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/pkg/app_error"
	"server/pkg/custom_error"
)

type comparer struct {
	redis      *redis.Client
	codeSaver  *saver
	key        string
	redisValue string
	code       string
	attempts   int
}

func newComparer(redis *redis.Client, codeSaver *saver) *comparer {
	return &comparer{
		redis:     redis,
		codeSaver: codeSaver,
	}
}

func (cc *comparer) CompareCodesAndIncrementOrDeleteIfNotEqual(key string, userCode string, maxAttempts int) error {
	cc.key = key

	err := cc.loadRedisValue()
	if err != nil {
		return err
	}

	err = cc.parseRedisValue()
	if err != nil {
		return err
	}

	if cc.compareCodesAndDeleteIfEqual(userCode) {
		return nil
	}

	err = cc.deleteCodeOrIncrementIfAttemptsLessThen(maxAttempts - 1)
	if err != nil {
		return err
	}

	if cc.deleted() {
		return custom_error.ErrCodeMaxAttempts
	}

	return custom_error.ErrCodesNotEqual
}

func (cc *comparer) loadRedisValue() error {
	redisValue, err := cc.redis.Get(context.TODO(), cc.key).Result()

	if errors.Is(err, redis.Nil) {
		return custom_error.ErrCodeIsNotInRedis
	}

	if err != nil {
		return app_error.New(err)
	}

	cc.redisValue = redisValue
	return nil
}

func (cc *comparer) parseRedisValue() error {
	_, err := fmt.Sscanf(cc.redisValue, codeFormat, &cc.code, &cc.attempts)
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (cc *comparer) compareCodesAndDeleteIfEqual(userCode string) bool {
	if cc.code != userCode {
		return false
	}

	go func() {
		_ = cc.deleteCode()
	}()

	return true
}

func (cc *comparer) deleteCodeOrIncrementIfAttemptsLessThen(maxAttempts int) error {
	if cc.attemptsLessThen(maxAttempts) {
		if err := cc.incrementAttempts(); err != nil {
			return app_error.New(err)
		}
		return nil
	}

	if err := cc.deleteCode(); err != nil {
		return app_error.New(err)
	}
	return nil
}

func (cc *comparer) attemptsLessThen(maxAttempts int) bool {
	return cc.attempts < maxAttempts
}

func (cc *comparer) incrementAttempts() error {
	ttl, err := cc.redis.TTL(context.TODO(), cc.key).Result()
	if err != nil {
		return app_error.New(err)
	}

	cc.attempts += 1

	return cc.codeSaver.Save(cc.key, cc.code, cc.attempts, ttl)
}

func (cc *comparer) deleteCode() error {
	cc.code = ""
	cc.redisValue = ""

	err := cc.redis.Del(context.TODO(), cc.key).Err()
	if err != nil {
		return app_error.New(err)
	}

	return nil
}

func (cc *comparer) deleted() bool {
	return cc.code == ""
}
