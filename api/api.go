package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilhaamms/user-management-api/controller"
)

type API struct {
	userController controller.UserController
}

func NewAPI(userController controller.UserController) *API {
	return &API{userController}
}

func (a *API) RegisterRoutes() *mux.Router {
	mux := mux.NewRouter()

	api := mux.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/user/register", a.userController.Register).Methods(http.MethodPost)

	return mux
}

func (a *API) Start() {
	mux := a.RegisterRoutes()

	log.Println("Server started at :8080")

	http.ListenAndServe(":8080", mux)
}
