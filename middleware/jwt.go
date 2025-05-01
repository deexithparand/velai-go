package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// verify the token with the signature
func JWTMiddleware(c *fiber.Ctx) error {

	log.Info("Entered Middleware")

	// authHeader := c.Get("Authorization")

	// if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
	// 	return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
	// }

	// tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	// token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
	// 	return utils.JwtSecret, nil
	// })

	// if err != nil || !token.Valid {
	// 	return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	// }

	// // Optionally set user info in context
	// claims := token.Claims.(jwt.MapClaims)
	// c.Locals("username", claims["username"])

	return c.Next()
}
