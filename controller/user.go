package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ilhaamms/user-management-api/helper"
	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/service"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	user := request.UserRegister{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := c.userService.Register(user)
	if err != nil {
		helper.ResponseJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseJsonSuccess(w, http.StatusCreated, "User has been registered", data)
}
