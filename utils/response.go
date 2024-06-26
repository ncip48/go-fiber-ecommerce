package utils

import "github.com/gofiber/fiber/v2"

func ResponseJson(c *fiber.Ctx, success bool, message any, data any) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": success,
		"msg":     message,
		"data":    data,
	})
}
