package services

import (
	"user-management-api/internal/models"
	"user-management-api/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetUserByUUID(uuid string) (*models.User, error) {
	return s.repo.GetByID(uuid)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) UpdateUser(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) DeleteUser(uuid string) error {
	return s.repo.Delete(uuid)
}
