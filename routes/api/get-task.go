package api

import (
	"context"
	"jwttest/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Task represents the task structure from the database
type Task struct {
	ID       string `json:"id" bson:"id"`
	Title    string `json:"title" bson:"title"`
	Status   string `json:"status" bson:"status"`
	Label    string `json:"label" bson:"label"`
	Priority string `json:"priority" bson:"priority"`
	Agent    string `json:"agent" bson:"agent"`
}

// RequestBody to extract email from the incoming request
type RequestBody struct {
	Email string `json:"email"`
}

func GetTasks(c *fiber.Ctx) error {
	// Parse JSON body to extract the email
	var body RequestBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	// MongoDB context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find tasks with matching agent email
	filter := bson.M{"agent": body.Email}
	cursor, err := db.TaskCollection.Find(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}
	defer cursor.Close(ctx)

	var tasks []Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode tasks",
		})
	}

	return c.JSON(tasks)
}
