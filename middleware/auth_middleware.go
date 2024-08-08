package middleware

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(AuthMiddleware)
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])

	return c.Next()
}
