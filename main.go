package main

import (
	"jwttest/routes"
	"jwttest/routes/auth"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Server Started")

	app := fiber.New()

	// health check endpoint
	app.Get("/health", routes.HealthCheck)

	// signup endpoint
	app.Post("/signup", auth.Signup)

	// login endpoint
	app.Post("/login", auth.Login)

	// protected := app.Group("/api")

	app.Listen("127.0.0.1:8000")
}
