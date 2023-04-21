package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"

	"github.com/labstack/echo/v4"
)

func ActivityRoutes(e *echo.Group) {
	activityRepository := repositories.RepositoryActivity(mysql.DB)
	h := handlers.HandlerActivity(activityRepository)

	e.GET("/activity", h.GetAllActivity)
	e.GET("/activity/:id", h.GetOneActivity)

	e.POST("/activity", h.CreateActivity)
	e.PATCH("/activity/:id", h.UpdateActivity)
	e.DELETE("/activity/:id", h.DeleteActivity)
}