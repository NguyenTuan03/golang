package handler

import (
	"gin/utils"
	"net/http"
	"regexp"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct{}

func NewProductHandler() *productHandler {
	return &productHandler{}
}

func (product *productHandler) GetListProducts(ctx *gin.Context) {
	searchStr := ctx.Query("search")
	limitStr := ctx.DefaultQuery("limit", "10")

	if err := utils.ValidationRequired("Search", searchStr); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	limitInt, err := strconv.Atoi(limitStr)
	if err != nil || limitInt < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Limit mus be positive",
		})
		return
	}

	searchRegex := regexp.MustCompile(`^[a-zA-Z0-9\s]{3,50}$`)
	if !searchRegex.MatchString(searchStr) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "liu liu",
			"message": "Search is not empty, only 3-50 chars, only character and white space",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List products",
		"search":  searchStr,
		"limit":   limitInt,
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

func (product *productHandler) GetProductionBySlug(ctx *gin.Context) {
	slugStr := ctx.Param("slug")
	//rule
	re := regexp.MustCompile("^[a-zA-Z0-9.-]*$")

	result := re.MatchString(slugStr)

	if result == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "product slug bad request",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "product slug",
		"product": slugStr,
	})
}

func (product *productHandler) GetProductsByCategory(ctx *gin.Context) {
	cate := ctx.Param("category")

	slice := make([]string, 3, 5)
	slice = append(slice, "php", "python", "golang")

	if !slices.Contains(slice, cate) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The category you gave does not have that subject",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category is...",
		"product": cate,
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
