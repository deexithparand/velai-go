package api

import "github.com/gofiber/fiber/v2"

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordhash"`
}

// list users from DB and send response back
func GetUsers(c *fiber.Ctx) error {

	//

	return nil
}

// add users to DB and send response back
func AddUsers(c *fiber.Ctx) error {

	return nil
}
