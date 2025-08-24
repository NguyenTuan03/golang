package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItem struct {
	ID          int        `json:"id" gorm:"column:id;"`
	Title       string     `json:"title" gorm:"column:title;"`
	Description string     `json:"description" gorm:"column:description;"`
	Status      string     `json:"status" gorm:"column:status;"`
	CreatedAt   *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	Status      string `json:"status" gorm:"column:status;"`
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Pagination) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("", GetItems(db))
			items.GET("/:id", GetItemById(db))
			items.PUT("/:id", UpdateItem(db))
			items.DELETE("/:id", DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreate
		data.Status = "pending"
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": data.Id,
		})
	}
}
func GetItems(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging Pagination
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		paging.Process()

		var result []TodoItem
		db = db.Where("COALESCE(status, '') <> ?", "Deleted")
		if err := db.Count(&paging.Total).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Order("id asc").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&result).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]any{
				"items": result,
				"page":  paging,
			},
			"message": "Get items successful",
			"status":  http.StatusOK,
		})
	}
}
func GetItemById(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		itemid := ctx.Param("id")

		var data TodoItem
		id, err := strconv.Atoi(itemid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
func UpdateItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var updatedData TodoItemUpdate
		itemid := ctx.Param("id")

		id, err := strconv.Atoi(itemid) // ID đã tìm

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err := ctx.ShouldBind(&updatedData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Tạo pointer cho status
		completedStatus := "completed"
		updatedData.Status = &completedStatus
		fmt.Println(updatedData)
		if err := db.Where("id = ?", id).Updates(&updatedData).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "item updated successful",
			"data":    updatedData,
		})
	}
}
func DeleteItem(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		itemid := ctx.Param("id")

		// var data TodoItem
		id, err := strconv.Atoi(itemid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Deleted successful",
		})
	}
}
