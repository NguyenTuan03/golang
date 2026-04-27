package repository

import (
	"user-management-api/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(uuid string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "uuid = ?", uuid).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll(page, limit int, search string) ([]models.User, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	var users []models.User
	offset := (page - 1) * limit
	if err := r.db.Where("name LIKE ?", "%"+search+"%").Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Where("uuid = ?", user.UUID).Updates(user).Error
}

func (r *userRepository) Delete(uuid string) error {
	return r.db.Model(&models.User{}).Where("uuid = ?", uuid).Update("status", models.UserStatusInactive).Error
}
