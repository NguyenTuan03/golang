package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID      string         `gorm:"primaryKey;type:varchar(36)" json:"uuid"`
	Name      string         `json:"name"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Age       int            `json:"age"`
	Password  string         `json:"password"`
	Status    UserStatus     `gorm:"type:string" json:"status"`
	Level     UserLevel      `gorm:"type:string" json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
)

type UserLevel string

const (
	UserLevelAdmin UserLevel = "ADMIN"
	UserLevelUser  UserLevel = "USER"
)
