package models

type User struct {
	UUID string `json:"uuid"`
	Name     string `json:"name"`
	Email string `json:"email"`
	Age  int    `json:"age"`
	Password string `json:"password"`
	Status UserStatus `json:"status"`
	Level UserLevel `json:"level"`
}

type UserStatus string

const (
	UserStatusActive UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
)

type UserLevel string

const (
	UserLevelAdmin UserLevel = "ADMIN"
	UserLevelUser UserLevel = "USER"
)