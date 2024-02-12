package handlers

import (
	"github.com/gofiber/fiber/v2"

	"tap/config"
	v "tap/internal/libs/validate"
	m "tap/internal/models"
)




func (h *Handler) SendEmail( c *fiber.Ctx) error {
	var email m.EmailRequest
	if err := c.BodyParser(&email); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}
	if err := v.Struct(email); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := h.service.SendEmail(email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Email sent",
	})

}

func (h *Handler) VerifyEmail( c *fiber.Ctx) error {
	verifyCode := c.Params("id")
	redirect := c.Query("redirect")
	err := h.service.VerifyEmail(verifyCode)
	if err != nil {
		return c.Redirect(config.ClientHome(), fiber.StatusMovedPermanently)
	}
	return c.Redirect(redirect, fiber.StatusMovedPermanently)
	
}