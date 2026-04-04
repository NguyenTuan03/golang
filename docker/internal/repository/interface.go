package repository

import "docker/internal/models"

type IUserRepository interface {
	GetListUsers()
	CreateUser(user *models.User)
	GetUserById(id int)
}