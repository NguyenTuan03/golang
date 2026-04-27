package constants

type EnvKey string

const (
	DBHostKey     EnvKey = "DB_HOST"
	DBPortKey     EnvKey = "DB_PORT"
	DBNameKey     EnvKey = "DB_NAME"
	DBUserKey     EnvKey = "DB_USER"
	DBPasswordKey EnvKey = "DB_PASSWORD"
	DBSSLModeKey  EnvKey = "DB_SSLMODE"
	PortKey       EnvKey = "PORT"
)

func (e EnvKey) String() string {
	return string(e)
}
