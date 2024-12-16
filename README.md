### Создание миграций

- Установить глобально пакет: `go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- Создание миграции: `migrate create -ext sql -dir migrations 'name'`