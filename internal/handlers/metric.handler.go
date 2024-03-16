package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"

	m "tap/internal/models"
)

func (h *Handler) VisitMetric(ctx *fiber.Ctx) error {
	pageId := ctx.Params("pageId")
	metric := m.Metric{
		PageID: pageId,
		Type:   "visit",
		Value:  "1",
		Date:   time.Now(),
	}
	err := h.service.CreateMetric(metric)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

func (h *Handler) UserMetric(ctx *fiber.Ctx) error {
	pageId := ctx.Params("pageId")
	metric := m.Metric{
		PageID: pageId,
		Type:   "user",
		Value:  "1",
		Date:   time.Now(),
	}
	err := h.service.CreateMetric(metric)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}

type ClickRequest struct {
	Value string `json:"value"`
}

func (h *Handler) ClickMetric(ctx *fiber.Ctx) error {
	pageId := ctx.Params("pageId")
	var click ClickRequest
	if err := ctx.BodyParser(&click); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		})
	}
	metric := m.Metric{
		PageID: pageId,
		Type:   "click",
		Value:  click.Value,
		Date:   time.Now(),
	}
	err := h.service.CreateMetric(metric)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
