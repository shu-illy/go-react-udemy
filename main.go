package main

import (
	"go-react-udemy/controller"
	"go-react-udemy/db"
	"go-react-udemy/repository"
	"go-react-udemy/router"
	"go-react-udemy/usecase"
	"go-react-udemy/validator"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	taskValidator := validator.NewTaskValidator()
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start("localhost:8080"))
}
