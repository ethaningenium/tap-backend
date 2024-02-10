package handlers

import (
	m "tap/internal/models"

	"github.com/gofiber/fiber/v2"
)


func (h *Handler) GetPage(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid address",
		})
	}
	page, err := h.service.GetPageByAddress(address)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(page)
}

func (h *Handler) CreatePage(c *fiber.Ctx) error {
	var page m.PageFromBody
	if err := c.BodyParser(&page); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}
	userId := c.Locals("user_id").(string)

	err := h.service.CreatePage(page, userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	return c.JSON(page)
}

func (h *Handler) UpdatePage(c *fiber.Ctx) error {
	var page m.PageFromBody
	if err := c.BodyParser(&page); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}

	userId := c.Locals("user_id").(string)
	err := h.service.UpdatePage(page, userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(page)

}