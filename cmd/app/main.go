package main

import (
	"log"
	"personcrud/internal/app"
)

func main() {

	application, err := app.New()
	if err != nil {
		log.Fatal("Failed to create app:", err)
	}
	defer application.Close()

	if err := application.Run(); err != nil {
		log.Fatal("Failed to run app:", err)
	}
}
