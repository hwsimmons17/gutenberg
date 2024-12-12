package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) AttachUsersRoutes() {
	a.engine.GET("/user", func(c *gin.Context) {
		_, err := c.Cookie("user_id")
		if err == nil {
			log.Println("user already has a cookie")
			return
		}
		if err != http.ErrNoCookie {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		userID := uuid.NewString()

		c.SetCookie("user_id", userID, 2147483647, "/", "", true, false)
		c.JSON(200, gin.H{"status": "ok"})
	})
}

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr, err := c.Cookie("user_id")
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized, hit `/user` first to get a user_id via cookies"})
			return
		}
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(500, gin.H{"error": "error parsing user_id"})
			return
		}
		c.Set("user_id", userID)
	}
}
