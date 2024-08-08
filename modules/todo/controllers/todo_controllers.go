package controllers

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/models"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/services"
	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	todoService services.TodoService
}

func NewTodoController(todoService services.TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (c *TodoController) GetTodos(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	todos, err := c.todoService.GetTodos(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(todos)
}

func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	todo := new(models.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	todo.UserID = userID
	if err := c.todoService.CreateTodo(todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(todo)
}

func (c *TodoController) GetTodoByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	todo, err := c.todoService.GetTodoByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(todo)
}

func (c *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	todo := new(models.Todo)
	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	todo.ID = uint(id)
	if err := c.todoService.UpdateTodo(todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(todo)
}

func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	if err := c.todoService.DeleteTodo(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
