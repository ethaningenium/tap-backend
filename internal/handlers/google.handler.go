package handlers

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	"tap/config"
)


func (h *Handler) AuthGoogle (ctx *fiber.Ctx) error {
	url := config.GoogleConfig().AuthCodeURL(config.GoogleState())
	return ctx.Redirect(url, fiber.StatusMovedPermanently)
}

func (h *Handler) AuthGoogleCallback (ctx *fiber.Ctx) error {
	code := ctx.Query("code")
	c := context.Background()
	token, err := config.GoogleConfig().Exchange(c, code)
	if err != nil {
		log.Println("Failed to exchange token:", err)
		return ctx.Redirect(config.ClientHome(), fiber.StatusMovedPermanently)
	}
	_ , refreshToken, err := h.service.AuthGoogle(token.AccessToken)
	ctx.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})
	return ctx.Redirect(config.ClientHome(), fiber.StatusMovedPermanently)
}

