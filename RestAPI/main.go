package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at,omitempty"`
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
	now := time.Now().UTC()
	item := TodoItem{
		Id:          1,
		Title:       "This is item 1",
		Description: "desc 1",
		Status:      "Doing",
		Created_at:  &now,
		Updated_at:  nil,
	}
	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v %T\n", string(jsonData), string(jsonData))
	jsonStr := "{\"id\":1,\"title\":\"This is item 1\",\"description\":\"desc 1\",\"status\":\"Doing\",\"created_at\":\"2025-08-09T14:15:37.7505723Z\",\"updated_at\":null}"

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v, %T", item2, item2)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item2,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
