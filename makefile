# Конфигурация базы данных
DB_HOST := localhost
DB_PORT := 5432
DB_NAME := postgres
DB_USER := postgres
DB_PASS := yourpassword
DB_DSN := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

# Команда migrate
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Команда psql внутри Docker-контейнера
PSQL := docker exec postgres-container psql -U $(DB_USER) -d $(DB_NAME)

# Цель по умолчанию
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make migrate        - Apply all migrations"
	@echo "  make migrate-down   - Roll back last migration"
	@echo "  make check-db       - List all tables in DB"
	@echo "  make run            - Run tasks gRPC server"
	@echo "  make clean          - Remove generated files (if any)"

# Применить миграции
migrate:
	@echo "Applying migrations to $(DB_HOST):$(DB_PORT) ..."
	$(MIGRATE) up

# Откатить последнюю миграцию
migrate-down:
	@echo "Rolling back last migration..."
	$(MIGRATE) down

# Проверить подключение к БД и посмотреть таблицы
check-db:
	@echo "Checking database tables..."
	$(PSQL) -c "\dt"

# Запуск сервера
run:
	@echo "Starting tasks gRPC server on :50052 ..."
	go run cmd/server/main.go

# Очистка (пока нет файлов для очистки)
clean:
	@echo "Nothing to clean (yet)."