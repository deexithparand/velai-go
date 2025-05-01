package utils

import (
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt/v5"
)

// just send a userID & get your token via a cookie
func GenerateJWT(userId string) (fiber.Cookie, error) {

	var secretKey string = "get me from the env file"
	var cookie fiber.Cookie

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
