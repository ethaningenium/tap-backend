package handlers

import (
	v "tap/internal/libs/validate"
	m "tap/internal/models"

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

	token, err := h.service.Register(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"user": user,
		})
	}
	c.Set("Authorization", token)


	return c.JSON(fiber.Map{
		"email": user.Email,
		"name": user.Name,
		"Authorization": token,
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


	name, token, err := h.service.Login(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}




	c.Set("Authorization", token)


	return c.JSON(fiber.Map{
		"email": user.Email,
		"name": name,
		"Authorization": token,
	})
}

func (h *Handler) Getme( c *fiber.Ctx) error {
	userId, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	if userId == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	user, err := h.service.Getme(userId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}
	return c.JSON(&fiber.Map{
		"email": user.Email,
		"name": user.Name,
		"isemailverified": user.IsEmailVerified,
	})
}

