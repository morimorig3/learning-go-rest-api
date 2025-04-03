package controller

import (
	"learning-go-rest-api/model"
	"learning-go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUseCase
}

func NewTaskController(tu usecase.ITaskUseCase) ITaskController {
	return &taskController{tu: tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// userから送られてくる jwtトークンからuser_idを取り出す
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64) // コンテキストから取得した値は any型 になるので型アサーションする

	tasksRes, err := tc.tu.GetAllTasks(uint(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	taskRes, err := tc.tu.GetTaskById(uint(userId), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId)
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	taskRes, err := tc.tu.UpdateTask(task, uint(userId), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(float64)

	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := tc.tu.DeleteTask(uint(userId), uint(taskId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
