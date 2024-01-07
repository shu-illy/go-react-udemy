package main

import (
	"go-react-udemy/controller"
	"go-react-udemy/db"
	"go-react-udemy/repository"
	"go-react-udemy/router"
	"go-react-udemy/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start("localhost:8080"))
}
