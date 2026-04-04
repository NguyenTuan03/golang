package config

import (
	"docker/internal/utils"
	"fmt"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Config struct {
	Database DatabaseConfig
}

func LoadConfig() (*Config, error) {
	return &Config{}, nil
}

func (c *Config) GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     utils.GetEnv("DB_HOST", "localhost"),
		Port:     utils.GetEnv("DB_PORT", "5433"),
		User:     utils.GetEnv("DB_USER", "postgres"),
		Password: utils.GetEnv("DB_PASSWORD", "postgres"),
		DBName:   utils.GetEnv("DB_NAME", "user_management_db"),
		SSLMode:  utils.GetEnv("DB_SSLMODE", "disable"),
	}
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf(
		PostgresConnectionStringTemplate,
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
