package auth

import (
	"encoding/json"
	"jwttest/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Usercreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// need to pass the username and password
func rehashnverify() error {
	// rehash using the same algo and check if the password matches
	// return errors.New("invalid creds")
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
