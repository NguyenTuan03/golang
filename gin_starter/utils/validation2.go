package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HanndleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, err := range validationError {
			switch err.Tag() {
			case "gt": 
			errors[err.Field()] = "Phai lon hon 0"
			}
		}
		return gin.H{
			"error": errors,
		}
	}
	return gin.H{
		"error": "Yeu cau khong hop le" + err.Error(),
	}
}