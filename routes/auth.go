package routes

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Usercreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// need to pass the username and password
func hashnload() error {
	// hash and load user in the database
	return nil
}

// need to pass the username and password
func rehashnverify() error {
	//rehash using the same algo and check if the password matches
	return errors.New("invalid creds")
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

	c.SendString("user created for " + usercreds.Username + " with id : " + " uid234")

	return nil
}

func Login(c *fiber.Ctx) error {
	var (
		usercreds Usercreds
		err       error
	)

	// post request with username and password
	requestBody := c.Body()

	log.Info("Request entry for login successful")

	// parsing the body
	err = json.Unmarshal(requestBody, &usercreds)
	if err != nil {
		log.Errorf("Error parsing the request body : ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}

	//rehashnverify
	err = rehashnverify()
	if err != nil {
		log.Errorf("password doesn't match ", err)
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid Credentials")
	}

	c.SendString("Successfully logged in")

	return nil
}
