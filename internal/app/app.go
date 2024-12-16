package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"server/internal/composites"
	pkgconfig "server/internal/config"
	"server/internal/constants"
	"server/migrations"
	"server/pkg/logger"
	pkgpostgres "server/pkg/postgres"
	pkgredis "server/pkg/redis"
	"time"
)

var ctx = context.Background()

func New() {
	// Load cfg
	cfg := pkgconfig.MustLoad("config/env.yaml")

	// Init logger
	logger.MustInitGlobal(constants.Environment(cfg.ENV))

	// Run postgres
	postgres, err := pkgpostgres.New(
		ctx,
		&pkgpostgres.Config{
			Host:                 cfg.Postgres.Host,
			Port:                 cfg.Postgres.Port,
			User:                 cfg.Postgres.User,
			Password:             cfg.Postgres.Password,
			DBName:               cfg.Postgres.DBName,
			SSLMode:              cfg.Postgres.SSLMode,
			ConnectionAttempts:   10,
			ConnectionSleepDelay: time.Second * 2,
		},
	)
	if err != nil {
		panic("Postgres not connected")
	}
	defer postgres.Close()

	// Run redis
	redis, err := pkgredis.New(
		ctx,
		&pkgredis.Config{
			Host:                 cfg.Redis.Host,
			Port:                 cfg.Redis.Port,
			Password:             cfg.Redis.Password,
			DB:                   cfg.Redis.DB,
			ConnectionAttempts:   10,
			ConnectionSleepDelay: time.Second * 2,
		},
	)
	if err != nil {
		panic("Redis not connected")
	}
	defer func() {
		_ = redis.Close()
	}()

	// Migrations
	migrations.Run(&migrations.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		User:     cfg.Postgres.User,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.DBName,
		SSLMode:  cfg.Postgres.SSLMode,
	})

	messageSender := composites.NewMessageSender(cfg)

	// Rest server
	restServer := echo.New()
	restServer.Logger.SetOutput(io.Discard)
	restApiGroup := restServer.Group("/api")

	composites.NewRestMiddlewares(restApiGroup)
	composites.NewAuth(cfg, restApiGroup, postgres, redis, messageSender)

	fmt.Println("Server started on port: " + cfg.Server.Port)
	if err = restServer.Start(":" + cfg.Server.Port); err != nil {
		panic(err)
	}
}
