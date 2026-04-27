package services

import (
	"errors"

	"user-management-api/internal/models"
	"user-management-api/internal/repository"

	"github.com/google/uuid"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("User name, email and password are required")
	}

	if user.Age <= 0 || user.Age > 120 {
		return errors.New("Invalid age")
	}

	_, err := s.repo.GetByEmail(user.Email)
	if err == nil {
		return errors.New("User with this email already exists")
	}

	user.UUID = uuid.New().String()
	user.Status = models.UserStatusActive
	user.Level = models.UserLevelUser

	return s.repo.Create(user)
}

func (s *userService) GetUserByUUID(uuid string) (*models.User, error) {
	return s.repo.GetByID(uuid)
}

func (s *userService) GetUsers(page, limit int, search string) ([]models.User, error) {
	return s.repo.GetAll(page, limit, search)
}

func (s *userService) UpdateUser(user *models.User) error {
	// Kiểm tra user có tồn tại không
	_, err := s.repo.GetByID(user.UUID)
	if err != nil {
		return errors.New("User not found")
	}

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(uuid string) error {
	// Check user exists
	_, err := s.repo.GetByID(uuid)
	if err != nil {
		return errors.New("User not found")
	}

	return s.repo.Delete(uuid)
}
