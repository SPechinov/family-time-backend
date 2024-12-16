package migrations

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"server/pkg/utils"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
}

func Run(options *Config) {
	fmt.Println("Applying migrations...")
	db, err := sql.Open(
		"postgres",
		utils.GetPostgresConnectionString(
			options.User,
			options.Password,
			options.Host,
			options.Port,
			options.DBName,
			options.SSLMode,
		),
	)
	if err != nil {
		panic("open: " + err.Error())
	}
	defer func() {
		_ = db.Close()
	}()

	// Инициализация драйвера миграций
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic("init: " + err.Error())
	}

	// Создание экземпляра миграции
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		panic("create: " + err.Error())
	}

	// Применение миграций
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic("apply: " + err.Error())
	}
	fmt.Println("Applied migrations")
}
