package main

import (
	"jwttest/db"
	"jwttest/middleware"
	"jwttest/routes"
	"jwttest/routes/api"
	"jwttest/routes/auth"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	// Load Environment Variables
	if os.Getenv("RENDER") == "" { // Only load .env in local dev
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		log.Println("Loaded environment variables from .env")
	}

	log.Println("Loaded environment variables")

	// Connect to database
	err := db.ConnectMongo(os.Getenv("MONGO_URI"), "velai-db")
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	}

	log.Println("Server Started")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://velai.onrender.com/", // Your frontend
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET,POST,OPTIONS",
		AllowCredentials: true,
	}))

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

	// route group for /api/users
	apiRouter.Get("/users", api.GetUsers)
	apiRouter.Post("/users", api.AddUsers)

	// route group for /api/tasks
	apiRouter.Post("/suggest", api.SuggestTasks)
	apiRouter.Post("/tasks", api.GetTasks)
	apiRouter.Post("/add-task", api.AddTask)
	apiRouter.Post("/delete-task", api.DeleteTask)

	app.Listen(":8000")
}
