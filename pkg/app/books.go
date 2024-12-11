package app

import (
	"gutenberg/pkg/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) AttachBooksRoutes() {
	a.engine.GET("/book/:id", getUser(), func(c *gin.Context) {
		user, ok := c.Get("user_id")
		if !ok {
			c.JSON(500, gin.H{"error": "error getting user_id"})
			return
		}
		bookID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{"error": "error parsing book id, must be an integer"})
			return
		}
		book, err := handlers.GetBook(c, bookID, user.(uuid.UUID), a.repository, a.bookReader)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, book)
	})

	a.engine.GET("/books", getUser(), func(c *gin.Context) {
		user, ok := c.Get("user_id")
		if !ok {
			c.JSON(500, gin.H{"error": "error getting user_id"})
			return
		}
		books, err := handlers.GetBooks(c, user.(uuid.UUID), a.repository)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, books)
	})
}
