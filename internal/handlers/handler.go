package handlers

import (
	m "tap/internal/models"
	"tap/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register( c *fiber.Ctx) error {
	var user m.UserRegister

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}

	refreshToken , accessToken, err := h.service.Register(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	c.Set("access_token", accessToken)

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"email": user.Email,
		"name": user.Name,
	})
}


func (h *Handler) Login( c *fiber.Ctx) error {
	var user m.UserLogin

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	name, refreshToken, accessToken, err := h.service.Login(user.Email, user.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	c.Set("access_token", accessToken)

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{
		"email": user.Email,
		"name": name,
	})
}