package middleware

import (
	"jwttest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt"
)

// verify the token with the signature
func JWTMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
		return c.Status(fiber.StatusUnauthorized).SendString("Missing or invalid Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	jwtSecret := utils.GetJWTSecret()
	log.Debugf("jwt secret retrieved : ", jwtSecret)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired token")
	}

	// Optionally set user info in context
	claims := token.Claims.(jwt.MapClaims)
	c.Locals("username", claims["username"])

	log.Info("Passed JWT Middleware")

	return c.Next()
}
