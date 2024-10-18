package repository

import (
	"database/sql"

	"github.com/ilhaamms/user-management-api/models/entity"
	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/models/response"
)

type UserRepository interface {
	Save(user request.UserRegister) (*response.UserRegisterResponse, error)
	FindByEmailRegister(email string) (response.User, error)
	FindByEmailLogin(email string) (*entity.User, error)
	FindAll() ([]response.User, error)
	UpdateById(id int, user request.UserUpdate) (*response.User, error)
	DeleteById(id int) (response.User, error)
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

	return &response.UserRegisterResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepository) FindByEmailRegister(email string) (response.User, error) {
	row := r.db.QueryRow("SELECT name, email FROM user WHERE email = ?", email)

	var user response.User
	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		return user, err
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

func (r *userRepository) FindAll() ([]response.User, error) {
	rows, err := r.db.Query("SELECT name, email FROM user")
	if err != nil {
		return nil, err
	}

	var users []response.User
	for rows.Next() {
		var user response.User
		err := rows.Scan(&user.Name, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) UpdateById(id int, user request.UserUpdate) (*response.User, error) {
	_, err := r.db.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
	if err != nil {
		return nil, err
	}

	return &response.User{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepository) DeleteById(id int) (response.User, error) {
	row := r.db.QueryRow("SELECT name, email FROM user WHERE id = ?", id)

	var user response.User
	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		return user, err
	}

	_, err = r.db.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return user, err
	}

	return user, nil
}
