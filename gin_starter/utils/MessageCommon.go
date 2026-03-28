package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *gin.Context, msg string, data any) {
	ctx.JSON(http.StatusOK, ResponseData{
		Code:    http.StatusOK,
		Message: msg,
		Data:    data,
	})
}

func ErrorResponse(ctx *gin.Context, statusCode int, errorCode ErrorCode, err error) {
	details := GetValidationErrorMsg(err)

	ctx.JSON(statusCode, ResponseData{
		Code:    int(errorCode),
		Message: errorCode.String(),
		Data:    details,
	})
}