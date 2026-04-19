package repository

import "user-management-api/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(uuid string) (*models.User, error)
	GetAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(uuid string) error
}
