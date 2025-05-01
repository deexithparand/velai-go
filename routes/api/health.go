package api

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error {
	c.SendString("Testing /api route")
	return c.Status(fiber.StatusOK).SendString("response from api route")
}
