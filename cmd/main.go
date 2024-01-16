package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"urlShorter/config"
	db "urlShorter/db/sqlc"
	"urlShorter/handlers"
	"urlShorter/usecase"
)

func main() {
	app := fiber.New()

	database := config.GetDBClient()

	databaseClient := db.New(database)
	urlUseCase := usecase.NewURL(databaseClient)
	encodeURL := handlers.NewEncodeURL(urlUseCase)

	app.Post("/shorten", encodeURL.Handler)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("error to start server, error %v", err.Error())
	}
}
