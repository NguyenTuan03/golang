package services

import "user-management-api/internal/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByUUID(uuid string) (*models.User, error)
	GetUsers() ([]models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(uuid string) error
}
