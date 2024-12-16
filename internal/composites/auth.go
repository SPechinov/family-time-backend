package composites

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	usersDatabase "server/internal/adapters/database/users"
	sessionsStore "server/internal/adapters/stores/sessions"
	authController "server/internal/api/rest/routes/auth"
	"server/internal/config"
	"server/internal/services/codes"
	"server/internal/services/message_sender"
	sessionsService "server/internal/services/sessions"
	usersService "server/internal/services/users"
	authUseCases "server/internal/usecases/auth"
)

func NewAuth(
	cfg *config.Config,
	router *echo.Group,
	database *pgxpool.Pool,
	redis *redis.Client,
	messageSender *message_sender.MessageSender,
) {
	db := usersDatabase.New(database)

	sStore := sessionsStore.New(redis)
	sService := sessionsService.New(cfg, sStore)

	us := usersService.New(cfg, db)
	cs := codes.New(redis)
	auc := authUseCases.New(cfg, us, cs, messageSender, sService)

	authController.New(cfg, auc).Register(router)
}
