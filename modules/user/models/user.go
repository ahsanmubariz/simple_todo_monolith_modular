package models

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/todo/models"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string        `json:"username" gorm:"unique"`
	Password string        `json:"-"`
	Todos    []models.Todo `json:"todos"`
}
