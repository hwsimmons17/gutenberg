package app

import (
	"gutenberg/pkg"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	engine     *gin.Engine
	repository pkg.BookRepository
	bookReader pkg.BookReader
}

func InitApp(
	repository pkg.BookRepository,
	bookReader pkg.BookReader,
) App {
	engine := gin.New()
	engine.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/drivers/location"),
		gin.Recovery(),
	)
	gin.SetMode(gin.ReleaseMode)

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://gutenberg-frontend.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Content-Disposition", "Sec-Websocket-Protocol"},
		ExposeHeaders:    []string{"Content-Length", "Content-Disposition"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	engine.Use(cors.New(config))

	return App{
		engine:     engine,
		repository: repository,
		bookReader: bookReader,
	}
}

func (a *App) Run() {
	a.AttachStandardRoutes()
	a.AttachUsersRoutes()
	a.AttachBooksRoutes()

	a.engine.Run()
}

func (a *App) AttachStandardRoutes() {
	a.engine.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
