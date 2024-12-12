package main

import (
	"gutenberg/pkg/app"
	"gutenberg/pkg/dependencies/claude"
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

	responseGenerator := claude.NewResponseGenerator(os.Getenv("CLAUDE_KEY"))

	server := app.InitApp(repository, bookReader, responseGenerator)

	server.Run()
}
