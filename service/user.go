package service

import (
	"errors"
	"math"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ilhaamms/user-management-api/models/entity"
	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/models/response"
	"github.com/ilhaamms/user-management-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user request.UserRegister) (*response.UserRegisterResponse, error)
	Login(user request.UserLogin) (*response.UserLoginResponse, error)
	FindAll(page, limit int) (*[]response.User, int, error)
	UpdateById(id int, user request.UserUpdate) (*response.User, error)
	DeleteById(id int) (*response.User, error)
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

	dbUser, _ := s.userRepository.FindByEmailRegister(user.Email)

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

	return &response.UserRegisterResponse{
		Name:  dataUSer.Name,
		Email: dataUSer.Email,
	}, nil
}

func (s *userService) Login(user request.UserLogin) (*response.UserLoginResponse, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("Email and Password are required")
	}

	dbUser, err := s.userRepository.FindByEmailLogin(user.Email)
	if err != nil {
		return nil, errors.New("Email or Password is incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return nil, errors.New("Email or Password is incorrect")
	}

	claims := &entity.Claims{
		Name: dbUser.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(entity.JwtKey))
	if err != nil {
		return nil, err
	}

	return &response.UserLoginResponse{
		Name:  dbUser.Name,
		Email: dbUser.Email,
		Token: tokenString,
	}, nil
}

func (s *userService) FindAll(page, limit int) (*[]response.User, int, error) {

	users, err := s.userRepository.FindAll()
	if err != nil {
		return nil, 0, err
	}

	if len(users) == 0 {
		return nil, 0, errors.New("users is empty")
	}

	startIndex := (page - 1) * limit
	endIndex := int(math.Min(float64(startIndex+limit), float64(len(users))))
	totalPages := int(math.Ceil(float64(len(users)) / float64(limit)))

	if page > totalPages {
		return nil, 0, errors.New("page sudah melebihi total page")
	}

	users = users[startIndex:endIndex]

	return &users, totalPages, nil
}

func (s *userService) UpdateById(id int, user request.UserUpdate) (*response.User, error) {

	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	if user.Name == "" || user.Email == "" {
		return nil, errors.New("Name and Email are required")
	}

	if len(user.Name) < 3 {
		return nil, errors.New("Name must be at least 3 characters")
	}

	if !strings.Contains(user.Email, "@") {
		return nil, errors.New("Email is not valid")
	}

	dataUser, err := s.userRepository.UpdateById(id, user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil, errors.New("Email already exists")
		} else {
			return nil, errors.New("User not found")
		}
	}

	return dataUser, nil
}

func (s *userService) DeleteById(id int) (*response.User, error) {
	if id <= 0 {
		return nil, errors.New("id must be greater than 0")
	}

	dataUser, err := s.userRepository.DeleteById(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return &dataUser, nil
}
