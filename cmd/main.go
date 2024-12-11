package main

import (
	"gutenberg/pkg/app"
	"gutenberg/pkg/dependencies/ebooks"
	"gutenberg/pkg/dependencies/postgres"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	postgresUrl := os.Getenv("DATABASE_URL")
	repository := postgres.InitRepository(postgresUrl)

	bookReader := ebooks.NewClient()

	server := app.InitApp(repository, bookReader)

	server.Run()
}
