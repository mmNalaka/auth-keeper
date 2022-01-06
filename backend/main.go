package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"github.com/mmnalaka/auth-keeper/app/router"
	"github.com/mmnalaka/auth-keeper/config"
	"github.com/mmnalaka/auth-keeper/database"
	"log"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Initialize
	config.Init()
	database.ConnectPostgres()

	// Initialize app
	app := fiber.New()

	// Middleware
	app.Use(requestid.New())
	app.Use(logger.New())

	// Register routes
	router.PublicRoutes(app)
	router.PrivateRoutes(app)
	router.NotFoundRouter(app)

	if err := app.Listen(config.Cfg.Server.Port); err != nil {
		log.Panic(err)
	}

	// Graceful shutdown
	// Listen from a different goroutine
	//go func() {
	//	if err := app.Listen(config.Cfg.Server.Port); err != nil {
	//		log.Panic(err)
	//	}
	//}()
	//
	//c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	//
	//_ = <-c // This blocks the main thread until an interrupt is received
	//fmt.Println("Gracefully shutting down...")
	//_ = app.Shutdown()
	//
	//fmt.Println("Running cleanup tasks...")
	//
	//// Cleanup tasks go here
	//fmt.Println("Fiber was successful shutdown.")
}
