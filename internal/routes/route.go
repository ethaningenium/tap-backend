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
	app.Get("/me",middle.Auth, handlers.Getme)
	app.Get("/page",middle.Auth, handlers.GetPages)
	app.Get("/page/:address", handlers.GetPage)
	app.Delete("/page/:address", middle.Auth, handlers.DeletePage)
	app.Post("/page", middle.Auth,  handlers.CreatePage)
	app.Put("/page", middle.Auth, handlers.UpdatePage)
	app.Patch("/page", middle.Auth, handlers.UpdateMeta)
	app.Get("/address/:check", handlers.CheckAddress)
	app.Get("/verify/:id", handlers.VerifyEmail)
	app.Post("/email", handlers.SendEmail)
	app.Get("/auth/google", handlers.AuthGoogle)
	app.Get("/auth/google/callback", handlers.AuthGoogleCallback)

	app.Get("/test", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "Token",
			Value:    "test",
			Expires:  time.Now().Add(time.Hour * 24 * 30),
			HTTPOnly: true,
		})
		c.Set("Authorization", "test")
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
	// TODO
}
