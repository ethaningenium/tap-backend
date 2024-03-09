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

	newId , err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	newPage := m.PageRequest{
		ID: page.ID,
		Title: page.Title,
		Address: page.Address,
		Theme: page.Theme,
		Favicon: page.Favicon,
		Bricks: page.Bricks,
		User: newId,
	}


	return c.JSON(newPage)
}

func (h *Handler) UpdatePage(c *fiber.Ctx) error {
	var page m.PageRequest
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
	if page.User != newId {
		return c.Status(500).JSON(fiber.Map{
			"message": "User ID does not match",
		})
	}
	err = h.service.UpdatePage(page)
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
	if page.User != newId {
		return c.Status(500).JSON(fiber.Map{
			"message": "User ID does not match",
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
	if userId == "" {
		return c.Status(500).JSON(fiber.Map{
			"message": "User ID not found",
		})
	}
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

func (h *Handler) DeletePage(c *fiber.Ctx) error {
	address := c.Params("address")
	userId := c.Locals("user_id").(string)
	err := h.service.DeletePage(address, userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}