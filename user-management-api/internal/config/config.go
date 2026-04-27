package config

import (
	"fmt"
	"log"
	"user-management-api/internal/constants"
	"user-management-api/internal/helper"

	"github.com/joho/godotenv"
)

type Config struct {
	DB  DatabaseConfig
	App AppConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	DSN      string
}

type AppConfig struct {
	Port string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default values")
	}

	// Database config
	dbHost := helper.GetEnv(constants.DBHostKey.String(), "localhost")
	dbPort := helper.GetEnv(constants.DBPortKey.String(), "5432")
	dbUser := helper.GetEnv(constants.DBUserKey.String(), "postgres")
	dbPass := helper.GetEnv(constants.DBPasswordKey.String(), "postgres")
	dbName := helper.GetEnv(constants.DBNameKey.String(), "user_management_db")
	dbSSL := helper.GetEnv(constants.DBSSLModeKey.String(), "disable")

	// Construct Postgres DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPass, dbName, dbPort, dbSSL)

	// App config
	appPort := helper.GetEnv(constants.PortKey.String(), "8080")
	if appPort[0] != ':' {
		appPort = ":" + appPort
	}

	return &Config{
		DB: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPass,
			DBName:   dbName,
			SSLMode:  dbSSL,
			DSN:      dsn,
		},
		App: AppConfig{
			Port: appPort,
		},
	}
}
