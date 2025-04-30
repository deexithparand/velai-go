package main

import (
	"jwttest/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Server Started")

	app := fiber.New()

	app.Get("/api/health", routes.GetHealthCheck)

	app.Listen("127.0.0.1:8000")
}
