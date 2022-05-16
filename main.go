package main

import (
	"github.com/gin-gonic/gin"

	"pustaka-api/book"
	"pustaka-api/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/db_pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&book.Book{})

	if err != nil {
		panic(err)
	}
	
	router := gin.Default()

	v1 := router.Group("v1")
	
	v1.GET("/", handler.RootHandler)

	v1.GET("/hello", handler.HelloHandler)

	v1.GET("/books/:id/:title", handler.BooksHandler)

	v1.GET("/query", handler.QueryHandler)

	v1.POST("/books", handler.PostBooksHandler)
	
	router.Run(":8888")
}
