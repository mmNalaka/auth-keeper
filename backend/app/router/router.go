package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mmnalaka/auth-keeper/app/controllers"
)

func handler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "pong",
	})
}

// PublicRoutes - Public Routes
func PublicRoutes(app *fiber.App) {
	// Create a new route group
	v1 := app.Group("/api/v1")

	// Register public routes
	v1.Post("/login", controllers.Login)
	v1.Post("/signup", controllers.Signup)

}

// PrivateRoutes - Private Routes
func PrivateRoutes(app *fiber.App) {
	// Create a new route group
	v1 := app.Group("/api/v1")

	// Register private routes

	v1.Get("/dashboard", monitor.New())
	v1.Get("/me", handler)
	v1.Get("/logout", handler)

}

// NotFoundRouter is a router for not 404 error
func NotFoundRouter(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Not Found",
		})
	})
}
