package todo

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/middleware"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/controllers"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/repositories"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/services"
	"github.com/gofiber/fiber/v2"
)

func SetupTodoRoutes(app *fiber.App) {
	todoRepository := repositories.NewTodoRepository()
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	api := app.Group("/api/todos", middleware.AuthMiddleware)
	api.Get("/", todoController.GetTodos)
	api.Post("/", todoController.CreateTodo)
	api.Get("/:id", todoController.GetTodoByID)
	api.Put("/:id", todoController.UpdateTodo)
	api.Delete("/:id", todoController.DeleteTodo)
}
