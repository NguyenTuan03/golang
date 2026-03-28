package handler2

import (
	"gin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)
type productHandler struct{}

func NewProductHandler() *productHandler {
	return &productHandler{}
}
type GetProductByIdRequest struct {
	ID int `uri:"id" binding:"gt=0"`
}
func (product *productHandler) GetProductById(ctx *gin.Context) {
	var params GetProductByIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HanndleValidationErrors(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Product ID",
		"product": params.ID,
	})
}