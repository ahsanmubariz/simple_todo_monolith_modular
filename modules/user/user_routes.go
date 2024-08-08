package user

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/controllers"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/repositories"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	api := app.Group("/api/users")
	api.Post("/register", userController.RegisterUser)
	api.Post("/login", userController.LoginUser)
}
