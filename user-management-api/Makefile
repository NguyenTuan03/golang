-include .env
export

# Connection string for migrations
CONN_STRING = postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_DIR = ./internal/database/migrations

# Docker commands
start:
	@echo "🚀 Đang khởi động các container Database..."
	docker compose up -d
	@echo "✅ Database đang chạy. Sử dụng 'make status' để kiểm tra."

stop:
	@echo "🛑 Đang dừng các container Database..."
	docker compose down

status:
	@echo "📊 Đang kiểm tra trạng thái các container..."
	docker compose ps

# Port management
clean-port:
	@echo "🧹 Đang giải phóng cổng 8386..."
	-lsof -ti:8386 | xargs kill -9 2>/dev/null || true

# Database operations
import-db:
	@echo "📥 Đang nhập dữ liệu vào database..."
	docker exec -i postgres-db psql -U $(DB_USER) -d $(DB_NAME) < ./backupdb-user-management-db.sql
	@echo "✅ Nhập dữ liệu hoàn tất."

export-db:
	@echo "📤 Đang xuất dữ liệu từ database..."
	docker exec -i postgres-db pg_dump -U $(DB_USER) -d $(DB_NAME) > ./backupdb-user-management-db.sql
	@echo "✅ Xuất dữ liệu hoàn tất tại ./backupdb-user-management-db.sql"

# Application commands
server: clean-port
	@echo "⚡ Đang khởi động Go API server..."
	go run cmd/api/main.go

# Migration commands
migrate-create:
	@echo "📝 Đang tạo file migration mới cho $(NAME)..."
	~/go/bin/migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)

migrate-up:
	@echo "🔼 Đang thực hiện migration up..."
	~/go/bin/migrate -path $(MIGRATIONS_DIR) -database "$(CONN_STRING)" up
	@echo "✅ Migration up hoàn tất."

migrate-down:
	@echo "🔽 Đang thực hiện migration down..."
	~/go/bin/migrate -path $(MIGRATIONS_DIR) -database "$(CONN_STRING)" down
	@echo "✅ Migration down hoàn tất."

.PHONY: import-db export-db server start stop status clean-port migrate-create migrate-up migrate-down