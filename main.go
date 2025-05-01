package main

import (
	"jwttest/middleware"
	"jwttest/routes"
	"jwttest/routes/api"
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

	apiRouter := app.Group("/api", middleware.JWTMiddleware)

	apiRouter.Get("/health", api.Health)

	app.Listen("127.0.0.1:8000")
}
