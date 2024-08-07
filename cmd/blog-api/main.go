package main

import "go-blog-api/internal/pkg/app"

func main() {
	application := app.NewApp()

	application.Run()
}
