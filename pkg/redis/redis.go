package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"server/pkg/utils"
	"time"
)

type Redis struct {
	client *redis.Client
}

type Config struct {
	Host                 string
	Port                 string
	Password             string
	DB                   int
	ConnectionAttempts   int
	ConnectionSleepDelay time.Duration
}

func New(ctx context.Context, config *Config) (*redis.Client, error) {
	var client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	err := connect(ctx, client, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func connect(ctx context.Context, client *redis.Client, config *Config) error {
	fmt.Println("Redis connecting...")

	err := utils.Repeatable(
		func() error {
			fmt.Println("Redis try to connect")

			pingErr := client.Ping(ctx).Err()
			if pingErr != nil {
				return pingErr
			}

			return nil
		},
		config.ConnectionAttempts,
		config.ConnectionSleepDelay,
	)

	if err != nil {
		fmt.Printf("Redis not connected: %s\n", err)
		return err
	}

	fmt.Println("Redis connected")
	return nil
}

func (r *Redis) Close() error {
	return r.client.Close()
}
