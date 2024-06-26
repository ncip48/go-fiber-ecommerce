package handlers

import (
	"ncip48/go-fiber-ecommerce/config"
	"ncip48/go-fiber-ecommerce/entities"
	"ncip48/go-fiber-ecommerce/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginAction(c *fiber.Ctx) error {
	login := &entities.Login{}

	// Parse JSON request body into login struct
	if err := c.BodyParser(login); err != nil {
		return utils.ResponseJson(c, false, "Invalid JSON", nil)
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(login); err != nil {
		// Return, if some fields are not valid.
		return utils.ResponseJson(c, false, utils.ValidatorErrors(err), nil)
	}

	user := entities.User{}

	result := config.Database.Last(&user, "username = ?", login.Username)

	if result.RowsAffected == 0 {
		return utils.ResponseJson(c, false, "User not found", nil)
	}

	comparePassword := utils.ComparePasswords(user.Password, login.Password)

	if !comparePassword {
		return utils.ResponseJson(c, false, "Wrong username or password", nil)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user": &user,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return utils.ResponseJson(c, true, "Success login", fiber.Map{
		"access_token": t,
	})
}

func GetProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	u := claims["user"]
	return utils.ResponseJson(c, true, "Success get profile", u)
}
