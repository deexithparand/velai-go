package api

import (
	"context"
	"jwttest/db"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTask(c *fiber.Ctx) error {
	// Extract Authorization header
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	_ = token // Optionally validate token here

	// Parse input from request body
	var body struct {
		Email  string `json:"email"`  // passed instead of agent
		TaskID string `json:"taskId"` // task-(digit)
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.Email == "" || body.TaskID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Both email and taskId are required",
		})
	}

	// MongoDB delete operation
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"agent": body.Email, // since agent is being set as email in AddTask
		"id":    body.TaskID,
	}

	result, err := db.TaskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete task",
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Task not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Task deleted successfully",
	})
}
