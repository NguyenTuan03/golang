package repository

import (
	"database/sql"
	"docker/internal/models"
	"fmt"
)

type SQLUserRepository struct {
	DB *sql.DB
}

func NewSQLUserRepository(db *sql.DB) IUserRepository {
	return &SQLUserRepository{DB: db}
}

func (r *SQLUserRepository) CreateUser(user *models.User) {
	
} 

func (r *SQLUserRepository) GetUserById(id int) {
	fmt.Println("user by id ", id)
}

func (r *SQLUserRepository) GetListUsers() {

}