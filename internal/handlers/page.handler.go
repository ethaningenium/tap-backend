package handlers

import (
	m "tap/internal/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (h *Handler) UpdateMeta(c *fiber.Ctx) error {
	var page m.PageMetaData
	if err := c.BodyParser(&page); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}

	userId := c.Locals("user_id").(string)
	newId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	page.User = newId
	err = h.service.UpdatePageMeta(page, userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(page)
}

func (h *Handler) GetPages(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(string)
	pages, err := h.service.GetPages(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(pages)
}

func (h *Handler) CheckAddress(c *fiber.Ctx) error {
	address := c.Params("check")

	exists, err := h.service.CheckAddress(address)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"exists": exists,
	})
}