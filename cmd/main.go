package main

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/robfig/cron.v2"
	db "urlShorter/db/sqlc"
	"urlShorter/infra/config"
	"urlShorter/infra/job"
	"urlShorter/internal/domain/encode"
	"urlShorter/internal/handlers"
	usecase2 "urlShorter/internal/usecase"
)

func main() {
	app := fiber.New()

	database := config.GetDBClient()
	encodeAdapter := usecase2.NewEncodeAdapter(base58.Encode)
	newEncode := encode.NewEncode(&encodeAdapter)

	databaseClient := db.New(database)
	urlUseCase := usecase2.NewURL(databaseClient)
	encodeURL := handlers.NewEncodeURL(&urlUseCase, &newEncode)

	cronExecute := cron.New()
	executer := job.NewExecuter(&urlUseCase, cronExecute)
	executer.Execute(context.Background())

	app.Post("/shorten", encodeURL.Create)
	app.Delete("/shorten/:id", encodeURL.Delete)
	app.Get("/:id", encodeURL.Get)

	err := app.Listen(":8080")
	if err != nil {
		fmt.Printf("error to start server, error %v", err.Error())
	}
}
