package main

import (
	"learning-go-rest-api/controller"
	"learning-go-rest-api/db"
	"learning-go-rest-api/repository"
	"learning-go-rest-api/router"
	"learning-go-rest-api/usecase"
	"learning-go-rest-api/validator"
)

func main() {
	db := db.NewDB()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	taskValidator := validator.NewTaskValidator()
	userValidator := validator.NewUserValidator()

	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
