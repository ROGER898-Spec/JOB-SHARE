package http

import (
	"github.com/FyaEdu/JOB-SHARE/backend/internal/delivery/http/handler"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/FyaEdu/JOB-SHARE/backend/docs"
)

func RegisterRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Jobshare Backend API is running smoothly!",
		})
	})

	api := app.Group("/api/v1")

	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
}
