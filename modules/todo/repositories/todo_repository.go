package repositories

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/config/database"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/models"
)

type TodoRepository interface {
	GetTodos(userID uint) ([]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id uint) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uint) error
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) GetTodos(userID uint) ([]models.Todo, error) {
	var todos []models.Todo
	if err := database.DB.Where("user_id = ?", userID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepository) CreateTodo(todo *models.Todo) error {
	return database.DB.Create(todo).Error
}

func (r *todoRepository) GetTodoByID(id uint) (*models.Todo, error) {
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todoRepository) UpdateTodo(todo *models.Todo) error {
	return database.DB.Save(todo).Error
}

func (r *todoRepository) DeleteTodo(id uint) error {
	var todo models.Todo
	return database.DB.Delete(&todo, id).Error
}
