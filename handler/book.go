package handler

import (
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"name" : "Ahmad Lamaul Farid",
		"bio" : "Devops Engineer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"title": "Hello World",
		"subtitle": "This is my first gin project",
	}) 
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H {
		"id" : id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) { 
	title := c.Query("title")
	subtitle := c.Query("subtitle")
	c.JSON(http.StatusOK, gin.H {
		"title": title,
		"subtitle": subtitle,
	})
}

func PostBooksHandler(c *gin.Context) {
	var book book.InputBook
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"title": book.Title,
		"price": book.Price,
		"author": book.Author,
	})
}