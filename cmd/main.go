package main

import (
	"gutenberg/pkg/app"
)

func main() {
	server := app.InitApp()

	server.Run()
}