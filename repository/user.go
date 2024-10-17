package repository

import (
	"database/sql"

	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/models/response"
)

type UserRepository interface {
	Save(user request.UserRegister) (*response.UserRegisterResponse, error)
	FindByEmail(email string) (response.UserRegisterResponse, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user request.UserRegister) (*response.UserRegisterResponse, error) {
	_, err := r.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	dataUSer := &response.UserRegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	return dataUSer, nil
}

func (r *userRepository) FindByEmail(email string) (response.UserRegisterResponse, error) {
	row := r.db.QueryRow("SELECT name, email FROM user WHERE email = ?", email)

	var user response.UserRegisterResponse
	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		return user, nil
	}

	return user, nil
}
