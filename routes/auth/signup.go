package auth

import (
	"context"
	"encoding/json"
	"jwttest/db"
	"jwttest/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// hash and load user into the database using bson.M
func hashnload(usercreds Usercreds) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usercreds.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf("Error hashing password: ", err)
		return err
	}

	// Create a bson.M object to represent the user document
	user := bson.M{
		"email":    usercreds.Email,
		"password": string(hashedPassword),
	}

	// Get MongoDB collection
	collection := db.UserCollection

	// Insert the user into the database
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Errorf("Error inserting user into database: ", err)
		return err
	}

	log.Info("User successfully inserted into MongoDB.")
	return nil
}

func Signup(c *fiber.Ctx) error {
	// usercreds object
	var (
		usercreds Usercreds
		err       error
	)

	log.Info("Request entry for signup successful")

	// get username and password from the request body
	requestBody := c.Body()

	// parse the string
	err = json.Unmarshal(requestBody, &usercreds)
	if err != nil {
		log.Errorf("Error parsing the request body : ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	// hash and store the password in the database
	log.Infof("user created and password stored in database...")
	err = hashnload(usercreds)
	if err != nil {
		log.Errorf("error occured while hashnload : ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	cookie, err := utils.GenerateJWT(usercreds.Email)
	if err != nil {
		log.Errorf("jwt error : ", err)
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
