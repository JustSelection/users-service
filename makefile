DB_HOST := localhost
DB_PORT := 5432
DB_NAME := postgres
DB_USER := postgres
DB_PASS := yourpassword
DB_DSN := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate:
	migrate -path ./migrations -database $(DB_DSN) up

migrate-down:
	migrate -path ./migrations -database $(DB_DSN) down