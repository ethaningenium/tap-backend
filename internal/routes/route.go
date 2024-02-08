package routes

import (
	"tap/internal/handlers"
	middle "tap/internal/middlewares"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func SetupRoutes (app *fiber.App, handlers *handlers.Handler) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Get("/me", handlers.Getme)
	app.Get("/test", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "refresh_token",
			Value:    "test",
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HTTPOnly: true,
		})
		return c.SendString("Hello, World!")
	})
	app.Get("/page/:address", handlers.GetPage)
	app.Post("/page",middle.CheckAuth,  handlers.CreatePage)
	// TODO
}
