package services

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/models"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/repositories"
)

type TodoService interface {
	GetTodos(userID uint) ([]models.Todo, error)
	CreateTodo(todo *models.Todo) error
	GetTodoByID(id uint) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uint) error
}

type todoService struct {
	todoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (s *todoService) GetTodos(userID uint) ([]models.Todo, error) {
	return s.todoRepository.GetTodos(userID)
}

func (s *todoService) CreateTodo(todo *models.Todo) error {
	return s.todoRepository.CreateTodo(todo)
}

func (s *todoService) GetTodoByID(id uint) (*models.Todo, error) {
	return s.todoRepository.GetTodoByID(id)
}

func (s *todoService) UpdateTodo(todo *models.Todo) error {
	return s.todoRepository.UpdateTodo(todo)
}

func (s *todoService) DeleteTodo(id uint) error {
	return s.todoRepository.DeleteTodo(id)
}
