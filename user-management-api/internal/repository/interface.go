package repository

import "user-management-api/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	GetByID(uuid string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll(page, limit int, search string) ([]models.User, error)
	Update(user *models.User) error
	Delete(uuid string) error
}
