package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"server/pkg/utils"
	"time"
)

type Config struct {
	Host                 string
	Port                 string
	User                 string
	Password             string
	DBName               string
	SSLMode              bool
	ConnectionAttempts   int
	ConnectionSleepDelay time.Duration
}

func New(ctx context.Context, config *Config) (*pgxpool.Pool, error) {
	pool, err := connect(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func connect(ctx context.Context, options *Config) (*pgxpool.Pool, error) {
	connectionString := utils.GetPostgresConnectionString(
		options.User,
		options.Password,
		options.Host,
		options.Port,
		options.DBName,
		options.SSLMode,
	)

	var pool *pgxpool.Pool

	fmt.Println("Postgres connecting...")
	err := utils.Repeatable(func() error {
		fmt.Println("Postgres try to connect")
		config, err := pgxpool.ParseConfig(connectionString)
		if err != nil {
			return err
		}

		config.MaxConns = 20
		config.MinConns = 5
		config.MaxConnIdleTime = 30 * time.Second

		pl, poolErr := pgxpool.NewWithConfig(ctx, config)
		if poolErr != nil {
			return poolErr
		}

		pingErr := pl.Ping(ctx)
		if pingErr != nil {
			return pingErr
		}

		pool = pl

		return nil
	},
		options.ConnectionAttempts,
		options.ConnectionSleepDelay,
	)

	if err != nil {
		fmt.Printf("Postgres not connected: %s\n", err)
		return nil, err
	}

	fmt.Println("Postgres connected")
	return pool, nil
}
