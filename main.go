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
	taskRepository := repository.NewTaskRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepository)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)

	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
