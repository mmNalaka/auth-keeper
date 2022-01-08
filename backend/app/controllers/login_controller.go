package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/auth-keeper/app/models"
	"github.com/mmnalaka/auth-keeper/app/queries"
	"github.com/mmnalaka/auth-keeper/config"
	"github.com/mmnalaka/auth-keeper/utils"
)

func Login(c *fiber.Ctx) error {

	// Login disabled
	if config.Cfg.App.DisableLogin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": "Login is disabled",
		})
	}

	// Get login params from request
	loginParams := &models.Login{}
	if err := c.BodyParser(loginParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	// Validate login params
	validate := utils.NewValidator()
	if err := validate.Struct(loginParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  utils.ValidatorErrors(err),
		})
	}

	// Find user
	user, err := queries.GetUserByEmail(loginParams.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"error":  utils.ValidatorErrors(err),
		})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"error":  "User not found, Invalid user name or password",
		})
	}

	// Check password
	if err := utils.CheckPasswordHash(loginParams.Password, user.Password); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"error":  "Invalid user name or password",
		})
	}

	return c.JSON(user)
}
