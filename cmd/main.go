package main

import (
	"github.com/gofiber/fiber/v2"
	"urlShorter/handlers"
)

func main() {

	app := fiber.New()

	app.Post("/shorte", handlers.Handler)

	app.Listen(":8080")

}
