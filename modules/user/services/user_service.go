package services

import (
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/models"
	"github.com/ahsanmubariz/simple_todo_monolith_modular/modules/user/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(username, password string) (*models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) RegisterUser(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return s.userRepository.CreateUser(user)
}

func (s *userService) LoginUser(username, password string) (*models.User, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
