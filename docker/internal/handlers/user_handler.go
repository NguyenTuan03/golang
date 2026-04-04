package handlers

import (
	"docker/internal/models"
	"docker/internal/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo repository.IUserRepository
}

func NewUserHandler(userRepo repository.IUserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (h *UserHandler) GetListUsers(c *gin.Context) {
	h.userRepo.GetListUsers()

	log.Println("users")

	c.JSON(http.StatusOK, gin.H{
		"message": "list users",
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	h.userRepo.GetUserById(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "user by id " + strconv.Itoa(id),
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	h.userRepo.CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
	})
}
