package handlers

import (
	"tap/internal/libs/jwt"
	m "tap/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register( c *fiber.Ctx) error {
	var user m.UserRegister

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}

	refreshToken , accessToken, err := h.service.Register(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"user": user,
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

func (h *Handler) Getme( c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	_ , err := jwt.VerifyToken(refreshToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	email, name, accessToken, err := h.service.Getme(refreshToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}
	c.Set("access_token", accessToken)
	return c.JSON(fiber.Map{
		"email": email,
		"name": name,
	})
}

