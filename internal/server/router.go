package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/kevinhartarto/workframe/internal/database"
)

func NewHandler(db database.Service) *fiber.App {
	context := context.Background()
	app := fiber.New()

	app.Use(healthcheck.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// API endpoints
	apiEndpoint := app.Group("/api")

	// user panel
	userAPI := apiEndpoint.Group("/user")
	userAPI.Post("/login", func(c *fiber.Ctx) error {
		return
	})

	userAPI.Post("/register", func(c *fiber.Ctx) error {
		return
	})

	// product panel
	productAPI := apiEndpoint.Group("/product")
}
