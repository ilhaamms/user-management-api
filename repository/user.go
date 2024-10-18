package repository

import (
	"database/sql"

	"github.com/ilhaamms/user-management-api/models/entity"
	"github.com/ilhaamms/user-management-api/models/request"
)

type UserRepository interface {
	Save(user request.UserRegister) (*entity.User, error)
	FindByEmailRegister(email string) (entity.User, error)
	FindByEmailLogin(email string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user request.UserRegister) (*entity.User, error) {
	_, err := r.db.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	dataUSer := &entity.User{
		Name:  user.Name,
		Email: user.Email,
	}

	return dataUSer, nil
}

func (r *userRepository) FindByEmailRegister(email string) (entity.User, error) {
	row := r.db.QueryRow("SELECT name, email FROM user WHERE email = ?", email)

	var user entity.User
	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (r *userRepository) FindByEmailLogin(email string) (*entity.User, error) {
	row := r.db.QueryRow("SELECT name, email, password FROM user WHERE email = ?", email)

	var user entity.User
	err := row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
