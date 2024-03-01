package main

import (
	"net/http"
	"task_teams/controller"
	"task_teams/database"
	"task_teams/exception"
	"task_teams/helper"
	"task_teams/repository"
	"task_teams/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Setup database connection
	// dialect := mysql.Open("root:@tcp(localhost:3306)/task1?charset=utf8mb4&parseTime=True&loc=Local")
	// db, err := gorm.Open(dialect, &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db := database.NewDB()
	db := database.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()
	router.POST("/api/users", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/logout", userController.Logout)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "Localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
