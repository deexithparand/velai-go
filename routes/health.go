package routes

import "github.com/gofiber/fiber/v2"

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("Hey Don't worry I'm Healthy!")
}
