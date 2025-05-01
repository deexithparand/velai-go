package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
)

func GetJWTSecret() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return secretKey
}

// just send a userID & get your token via a cookie
func GenerateJWT(userId string) (fiber.Cookie, error) {
	var (
		secretKey string
		cookie    fiber.Cookie
		err       error
	)

	secretKey = GetJWTSecret()

	// JWT Claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userId,
		"expiry": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return cookie, err
	}

	// Set JWT token in cookie
	cookie = fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24), // Expires in 24 hours
		HTTPOnly: true,
		Secure:   true,
	}

	return cookie, nil
}
