package middlewares

import (
	jwt "tap/internal/libs/jwt"

	"github.com/gofiber/fiber/v2"
)

func CheckAuth(c *fiber.Ctx) error {
	accessToken := c.Get("access_token")
	if accessToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	claims, err := jwt.VerifyToken(accessToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	c.Locals("user_id", claims.ID)
	return c.Next()
}