package main

import (
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gofiber/fiber/v2"
	"urlShorter/config"
	db "urlShorter/db/sqlc"
	"urlShorter/domain/encode"
	"urlShorter/handlers"
	"urlShorter/usecase"
)

func main() {
	app := fiber.New()

	database := config.GetDBClient()

	encodeAdapter := usecase.NewEncodeAdapter(base58.Encode)

	newEncode := encode.NewEncode(&encodeAdapter)

	databaseClient := db.New(database)
	urlUseCase := usecase.NewURL(databaseClient)
	encodeURL := handlers.NewEncodeURL(urlUseCase, newEncode)

	app.Post("/shorten", encodeURL.Create)
	app.Delete("/shorten/:id", encodeURL.Delete)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("error to start server, error %v", err.Error())
	}
}
