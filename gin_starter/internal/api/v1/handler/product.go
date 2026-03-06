package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct{}

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (product *productHandler) GetListProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List products",
	})
}

func (product *productHandler) GetProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	id, err := strconv.Atoi(productId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product id",
		"product": id,
	})
}

func (product *productHandler) CreateNewProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "creat a new product",
	})
}

func (product *productHandler) UpdateProductById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "update a product by id",
	})
}

func (product *productHandler) DeleteProductById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete product by id",
	})
}
