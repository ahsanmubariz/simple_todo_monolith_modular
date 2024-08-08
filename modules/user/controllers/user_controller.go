package controllers

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/models"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/services"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) RegisterUser(ctx *fiber.Ctx) error {
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.userService.RegisterUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(user)
}

func (c *UserController) LoginUser(ctx *fiber.Ctx) error {
	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dbUser, err := c.userService.LoginUser(user.Username, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	token, err := utils.GenerateJWT(dbUser.ID, dbUser.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
