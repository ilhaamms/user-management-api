package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ilhaamms/user-management-api/api"
	"github.com/ilhaamms/user-management-api/config"
	"github.com/ilhaamms/user-management-api/controller"
	"github.com/ilhaamms/user-management-api/repository"
	"github.com/ilhaamms/user-management-api/service"
)

func main() {

	db, err := config.GetConnection()
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	api := api.NewAPI(userController)
	api.Start()
}
