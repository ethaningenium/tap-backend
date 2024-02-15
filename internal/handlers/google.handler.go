package handlers

import (
	"context"
	"fmt"
	"log"

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
	googleToken, err := config.GoogleConfig().Exchange(c, code)
	if err != nil {
		log.Println("Failed to exchange token:", err)
		return ctx.Redirect(config.ClientHome(), fiber.StatusMovedPermanently)
	}
	token, err := h.service.AuthGoogle(googleToken.AccessToken)
	if err != nil {
		return ctx.Redirect(config.ClientHome(), fiber.StatusMovedPermanently)
	}
	
	return ctx.Redirect(fmt.Sprintf("%s?token=%s", config.ClientHome(), token), fiber.StatusMovedPermanently)
}

