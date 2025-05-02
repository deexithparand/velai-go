package api

import (
	"encoding/json"
	"jwttest/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// list users from DB and send response back
func GetUsers(c *fiber.Ctx) error {

	// get users from user collections
	cursor, err := db.UserCollection.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).SendString("Failed to fetch tasks")
	}
	defer cursor.Close(c.Context())

	var users []bson.M
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString("Error decoding tasks")
	}

	return c.JSON(users)
}

// add users to DB and send response back
func AddUsers(c *fiber.Ctx) error {

	var user User

	// POST request with the body having data in this format
	requestBody := c.Body()

	err := json.Unmarshal(requestBody, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// now save the user details to the database

	return nil
}
