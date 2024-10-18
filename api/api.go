package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhaamms/user-management-api/controller"
	"github.com/ilhaamms/user-management-api/middleware"
)

type API struct {
	userController controller.UserController
}

func NewAPI(userController controller.UserController) *API {
	return &API{userController}
}

func (a *API) RegisterRoutes() *mux.Router {
	mux := mux.NewRouter()

	auth := mux.PathPrefix("/api/v1/auth").Subrouter()
	auth.HandleFunc("/register", a.userController.Register).Methods(http.MethodPost)
	auth.HandleFunc("/login", a.userController.Login).Methods(http.MethodPost)

	users := mux.PathPrefix("/api/v1").Subrouter()
	users.Use(middleware.Auth)
	users.HandleFunc("/users", a.userController.GetAllUsers).Methods(http.MethodGet)
	users.HandleFunc("/users", a.userController.UpdateUser).Methods(http.MethodPatch)

	return mux
}

func (a *API) Start() {
	mux := a.RegisterRoutes()

	log.Println("Server started at :8080")

	http.ListenAndServe(":8080", mux)
}
