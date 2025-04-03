package main

import (
	"learning-go-rest-api/controller"
	"learning-go-rest-api/db"
	"learning-go-rest-api/repository"
	"learning-go-rest-api/router"
	"learning-go-rest-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
