package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"

	"github.com/labstack/echo/v4"
)

func ToDoRoutes(e *echo.Group) {
	todoRepository := repositories.RepositoryToDo(mysql.DB)
	h := handlers.HandlerToDo(todoRepository)

	e.GET("/todo", h.GetAllToDo)
	e.GET("/todo/:id", h.GetOneToDo)
	e.POST("/todo", h.CreateToDo)
	e.PATCH("/todo/:id", h.UpdateToDo)
	e.DELETE("/todo/:id", h.DeleteToDo)
}