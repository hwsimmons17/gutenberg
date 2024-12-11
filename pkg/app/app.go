package app

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


type App struct {
	engine                     *gin.Engine
}

func InitApp() App {
	engine := gin.New()
	engine.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/drivers/location"),
		gin.Recovery(),
	)
	gin.SetMode(gin.ReleaseMode)

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Content-Disposition", "Sec-Websocket-Protocol"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	engine.Use(cors.New(config))

	return App{
		engine: engine,
	}
}

func (a *App) Run() {
	a.AttachStandardRoutes()

	a.engine.Run()
}

func (a *App) AttachStandardRoutes() {
	a.engine.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}