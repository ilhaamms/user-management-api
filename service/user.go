package service

import (
	"errors"
	"strings"

	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/models/response"
	"github.com/ilhaamms/user-management-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user request.UserRegister) (*response.UserRegisterResponse, error)
	// Login(user request.UserLogin) (*response.UserLoginResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user request.UserRegister) (*response.UserRegisterResponse, error) {

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return nil, errors.New("Name, Email, and Password are required")
	}

	if len(user.Name) < 3 {
		return nil, errors.New("Name must be at least 3 characters")
	}

	if len(user.Password) < 6 {
		return nil, errors.New("Password must be at least 6 characters")
	}

	if !strings.Contains(user.Email, "@") {
		return nil, errors.New("Email is not valid")
	}

	dbUser, _ := s.userRepository.FindByEmail(user.Email)

	if dbUser.Email == user.Email {
		return nil, errors.New("Email already exists")
	}

	passwordBcrypt, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(passwordBcrypt)

	dataUSer, err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return dataUSer, nil
}

// func (s *userService) Login(user request.UserLogin) (*response.UserLoginResponse, error) {
// 	return s.userRepository.Login(user)
// }
