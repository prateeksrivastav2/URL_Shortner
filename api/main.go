package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/prateeksrivastav2/UrlShortner/api/routes"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1", routes.ShortenURL)
	app.Get("/:url", routes.ResolveURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
