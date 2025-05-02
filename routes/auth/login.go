package auth

import (
	"encoding/json"
	"errors"
	"jwttest/db"
	"jwttest/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type Usercreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// rehashnverify will retrieve the hashed password from the DB and verify it with the provided password
func rehashnverify(email, password string) (bool, error) {
	// Fetch user document from the database by email
	var user bson.M
	err := db.UserCollection.FindOne(db.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		// Return an error if no user is found or there's an issue with the query
		return false, errors.New("user not found or database error")
	}

	// Retrieve the hashed password from the user document
	hashedPassword, ok := user["password"].(string)
	if !ok {
		return false, errors.New("password field is missing or not a string")
	}

	// Compare the stored hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// If the password doesn't match, return an error
		return false, errors.New("invalid credentials")
	}

	// If the password matches
	return true, nil
}

func Login(c *fiber.Ctx) error {
	var (
		usercreds Usercreds
		err       error
	)

	// post request with username and password
	requestBody := c.Body()

	log.Info("Request entry for login successful")

	// Parsing the body
	err = json.Unmarshal(requestBody, &usercreds)
	if err != nil {
		log.Errorf("Error parsing the request body : ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// Call rehashnverify to check if the credentials are correct
	isValid, err := rehashnverify(usercreds.Email, usercreds.Password)
	if err != nil || !isValid {
		log.Errorf("Invalid credentials: %v", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// Generate JWT token if credentials are correct
	cookie, err := utils.GenerateJWT(usercreds.Email)
	if err != nil {
		log.Errorf("JWT error: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	// Set the JWT cookie in the response
	c.Cookie(&cookie)

	// Respond with the success message, including cookie (for debugging, or frontend handling)
	c.JSON(fiber.Map{
		"message": "User created successfully for " + usercreds.Email,
		"cookie":  cookie.Value, // Optionally, you can send back the cookie value if needed
	})

	return nil
}
