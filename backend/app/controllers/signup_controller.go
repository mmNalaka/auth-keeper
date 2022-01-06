package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mmnalaka/auth-keeper/app/models"
	"github.com/mmnalaka/auth-keeper/app/queries"
	"github.com/mmnalaka/auth-keeper/config"
	"github.com/mmnalaka/auth-keeper/utils"
)

func Signup(c *fiber.Ctx) error {

	// Login disabled
	if config.Cfg.App.DisableSignup {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  fiber.StatusForbidden,
			"message": "Signup is disabled",
		})
	}

	// Get signup params from request
	signupParams := &models.Signup{}
	if err := c.BodyParser(signupParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  err.Error(),
		})
	}

	// Validate signup params
	validate := utils.NewValidator()
	if err := validate.Struct(signupParams); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": fiber.StatusBadRequest,
			"error":  utils.ValidatorErrors(err),
		})
	}

	// Check if user already exists
	user, err := queries.GetUserByEmail(signupParams.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database error, finding user",
		})
	}

	if user != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"status": fiber.StatusConflict,
			"error":  "User already exists",
		})
	}

	user, err = queries.CreateUser(signupParams.Email, signupParams.Password, signupParams.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": fiber.StatusInternalServerError,
			"error":  "Database error, creating user",
		})
	}

	fmt.Printf("%+v\n", user)

	return c.JSON(fiber.Map{
		"status": fiber.StatusCreated,
		"data": fiber.Map{
			"user": user,
		},
	})
}
