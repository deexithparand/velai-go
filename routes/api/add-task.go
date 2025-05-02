package api

import (
	"context"
	"jwttest/db"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AddTask(c *fiber.Ctx) error {
	// Extract Authorization header
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	_ = token // You can validate token if needed

	// Parse task data from request body
	var task Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if task.Agent == "" || task.ID == "" || task.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Agent, ID, and Title fields are required",
		})
	}

	// Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert task",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task added successfully",
	})
}
