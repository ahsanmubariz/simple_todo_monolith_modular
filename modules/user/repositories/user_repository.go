package repositories

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/config/database"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
