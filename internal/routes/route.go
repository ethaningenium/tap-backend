package routes

import (
	"tap/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes (app *fiber.App, handlers *handlers.Handler) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	// TODO
}