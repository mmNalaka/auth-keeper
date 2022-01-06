package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/auth-keeper/app/models"
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

	return c.JSON(loginParams)
}
