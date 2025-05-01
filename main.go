package main

import (
	"jwttest/middleware"
	"jwttest/routes"
	"jwttest/routes/api"
	"jwttest/routes/auth"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Loaded environment variables")

	log.Println("Server Started")

	app := fiber.New()

	// health check endpoint
	app.Get("/health", routes.HealthCheck)

	// signup endpoint
	app.Post("/signup", auth.Signup)

	// login endpoint
	app.Post("/login", auth.Login)

	// api router routes
	apiRouter := app.Group("/api", middleware.JWTMiddleware)

	// health check for /api/health
	apiRouter.Get("/health", api.Health)

	// route group for /api/
	// apiRouter.Get("/task", api.task)

	app.Listen(":8000")
}
