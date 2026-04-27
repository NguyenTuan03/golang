package handler

import (
	"net/http"
	"strconv"

	"user-management-api/internal/models"
	"user-management-api/internal/services"
	"user-management-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	if err := h.service.CreateUser (&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Data:    user,
		Error:   "",
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	uuid := c.Param("uuid")
	user, err := h.service.GetUserByUUID(uuid)

	if err != nil {
		c.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Data:    user,
		Error:   "",
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	search := c.DefaultQuery("search", "")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	users, err := h.service.GetUsers(pageInt, limitInt, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Data:    users,
		Meta: utils.ResponseMeta{
			Page:  pageInt,
			Limit: limitInt,
			Total: len(users),
		},
		Error: "",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	user.UUID = c.Param("uuid")

	if err := h.service.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Data:    user,
		Error:   "",
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	uuidStr := c.Param("uuid")
	if _, err := uuid.Parse(uuidStr); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Data:    nil,
			Error:   "Invalid UUID format",
		})
		return
	}

	if err := h.service.DeleteUser(uuidStr); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Data:    nil,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Success: true,
		Data:    nil,
		Error:   "",
	})
}
