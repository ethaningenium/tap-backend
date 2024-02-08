package handlers

import (
	"tap/internal/libs/primitive"
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
	page := new(m.PageFromBody)
	if err := c.BodyParser(page); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}
	userId := c.Locals("user_id").(string)
	HexId, err := primitive.GetObject(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	requestPage := m.PageRequest{
		ID: page.ID,
		Title: page.Title,
		Address: page.Address,
		Bricks: page.Bricks,
		User: HexId,
	}

	err = h.service.CreatePage(requestPage)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	return c.JSON(fiber.Map{
		"message": "Hello, World!",
		"page": page,
	})
}