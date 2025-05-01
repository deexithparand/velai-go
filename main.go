package main

import (
	"jwttest/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Server Started")

	app := fiber.New()

	// health check endpoint
	app.Get("/api/health", routes.HealthCheck)

	// signup endpoint
	app.Post("/api/signup", routes.Signup)

	// login endpoint
	app.Post("/api/login", routes.Login)

	app.Listen("127.0.0.1:8000")
}
