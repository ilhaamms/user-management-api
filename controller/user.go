package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ilhaamms/user-management-api/helper"
	"github.com/ilhaamms/user-management-api/models/request"
	"github.com/ilhaamms/user-management-api/service"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
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

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	user := request.UserLogin{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, err := c.userService.Login(user)
	if err != nil {
		helper.ResponseJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	cookie := &http.Cookie{
		Name:    "name",
		Value:   data.Name,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, cookie)

	helper.ResponseJsonSuccess(w, http.StatusOK, "Login success", data)
}
