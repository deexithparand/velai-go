package auth

import (
	"encoding/json"
	"jwttest/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// need to pass the username and password
func hashnload() error {
	// hash and load user in the database
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
	err = hashnload()
	if err != nil {
		log.Errorf("error occured while hashnload : ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	// log and print
	log.Info("User created and password stored (mocked).")
	log.Infof("user created for %s with id %s", usercreds.Username, "uid234")

	cookie, err := utils.GenerateJWT("uid234")
	if err != nil {
		log.Errorf("jwt error : ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	c.JSON(fiber.Map{
		"message": "user created for " + usercreds.Username + " with id : " + " uid234",
		"cookie":  cookie,
	})

	return nil
}
