package database

import (
	"context"
	"database/sql"
	"docker/internal/config"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	dbConfig := cfg.GetDatabaseConfig()

	// Opening a driver typically will not attempt to connect to the database.
	DB, err = sql.Open(config.DriverPostgres, dbConfig.ConnectionString())
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		fmt.Println("unable to use data source name", err)
	}

	// Cấu hình Database Pool
	DB.SetConnMaxLifetime(30 * time.Minute)
	DB.SetMaxIdleConns(3)
	DB.SetMaxOpenConns(3)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	// Kiểm tra kết nối thực tế với Timeout
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	Ping(ctx)
	fmt.Println("connected to database")

	return nil
}

func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		fmt.Printf("unable to connect to database: %v", err)
	}
}
