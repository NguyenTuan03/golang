package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UUID      string         `gorm:"primaryKey;type:varchar(36)" json:"uuid"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"not null;uniqueIndex" json:"email"`
	Age       int            `gorm:"not null" json:"age"`
	Password  string         `gorm:"not null" json:"password"`
	Status    UserStatus     `gorm:"type:string;not null" json:"status"`
	Level     UserLevel      `gorm:"type:string;not null" json:"level"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// type UserStatus string
type UserStatus string

const (
	UserStatusActive  UserStatus = "ACTIVE"
	UserStatusInactive UserStatus = "INACTIVE"
)

type UserLevel string

const (
	UserLevelAdmin UserLevel = "ADMIN"
	UserLevelUser  UserLevel = "USER"
)

type RequestUUID struct {
	UUID string `uri:"uuid" binding:"required,uuid"` // Checks if 'id' is a valid UUID
}