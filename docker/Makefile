include .env
export DB_HOST

CONN_STRING = postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

DIR_STRING = ./internal/database/migrations

import-db:
	docker exec -i postgres-db psql -U postgres -d user_management_db < ./backupdb-user-management-db.sql

export-db:
	docker exec -i postgres-db pg_dump -U postgres -d user_management_db > ./backupdb-user-management-db.sql

server:
	go run .
	
start:
	docker-compose up -d

stop:
	docker-compose down

migrate-create:
	~/go/bin/migrate create -ext sql -dir $(DIR_STRING) -seq $(NAME)

migrate-up:
	~/go/bin/migrate -path $(DIR_STRING) -database "$(CONN_STRING)" up

migrate-down:
	~/go/bin/migrate -path $(DIR_STRING) -database "$(CONN_STRING)" down

.PHONY: import-db export-db server start stop migrate-create migrate-up migrate-down