package main

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/config/database"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/config/migrations"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/middleware"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	if err := database.ConnectDB(); err != nil {
		panic(err)
	}

	if err := migrations.RunMigrations(); err != nil {
		panic(err)
	}

	middleware.SetupMiddleware(app)
	todo.SetupTodoRoutes(app)
	user.SetupUserRoutes(app)

	app.Listen(":3000")
}
