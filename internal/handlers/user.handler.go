package handlers

import (
	"tap/internal/libs/jwt"
	v "tap/internal/libs/validate"
	m "tap/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register( c *fiber.Ctx) error {
	var user m.RegisterBody

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Error parsing body",
		}) 
	}
	if err := v.Struct(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
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
	var user m.LoginBody

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := v.Struct(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	name, refreshToken, accessToken, err := h.service.Login(user)
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
	claims , err := jwt.VerifyRefresh(refreshToken)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}
	name, accessToken, err := h.service.Getme(claims.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}
	c.Set("access_token", accessToken)
	return c.JSON(fiber.Map{
		"email": claims.Email,
		"name": name,
	})
}

