package utils

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// just send a userID & get your token via a cookie
func GenerateJWT(userId string) (fiber.Cookie, error) {

	var (
		secretKey string
		cookie    fiber.Cookie
		err       error
	)

	// GET jwt secret key
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey = os.Getenv("JWT_SECRET_KEY")

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
