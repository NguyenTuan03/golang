package handler2

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct{}

func NewUserHandler() *userHandler {
	return &userHandler{}
}

func (user *userHandler) GetListUsers(ctx *gin.Context) {
	log.Println("Get list users v2")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List users",
	})
}

func (user *userHandler) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("id")
	id, err := strconv.Atoi(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be higher or equal 1"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "users id",
		"user":    id,
	})
}

func (user *userHandler) GetUserByUuid(ctx *gin.Context) {
	idStr := ctx.Param("uuid")

	_,err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be uuid",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "users id",
		"user":    idStr,
	})
}

func (user *userHandler) CreateNewUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "creat a new user",
	})
}

func (user *userHandler) UpdateUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update a user by id",
	})
}

func (user *userHandler) DeleteUserById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete user by id",
	})
}
