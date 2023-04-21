package handlers

import (
	dto "BE/dto/result"
	tododto "BE/dto/todo"
	"BE/models"
	"BE/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerToDo struct {
	ToDoRepository repositories.ToDoRepository
}

func HandlerToDo(ToDoRepository repositories.ToDoRepository) *handlerToDo {
	return &handlerToDo{ToDoRepository}
}

func (h *handlerToDo) GetAllToDo(c echo.Context) error {
	ActivityID, _ := strconv.Atoi(c.QueryParam("activity_group_id"))
	
	ToDo, err := h.ToDoRepository.GetAllToDo(ActivityID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: ToDo})
}

func (h *handlerToDo) GetOneToDo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ToDo, err := h.ToDoRepository.GetOneToDo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: ToDo})
}

func (h *handlerToDo) CreateToDo(c echo.Context) error {
	request := new(tododto.ToDoRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	ToDo := models.ToDo{
		Title:     		request.Title,
		ActivityID:    	request.ActivityID,
		IsActive:		request.IsActive,
		Priority:		request.Priority,
		CreatedAt: 		time.Now(),
	}

	data, err := h.ToDoRepository.CreateToDo(ToDo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}

func (h *handlerToDo) UpdateToDo(c echo.Context) error {
	request := new(tododto.UpdateToDoRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	ToDo, err := h.ToDoRepository.GetOneToDo(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		ToDo.Title = request.Title
		ToDo.UpdateAt 	= time.Now()
	}

	if request.IsActive != ToDo.IsActive {
		ToDo.IsActive = request.IsActive
		ToDo.UpdateAt 	= time.Now()
	}

	if request.Priority != "" {
		ToDo.Priority = request.Priority
		ToDo.UpdateAt 	= time.Now()
	}

	data, err := h.ToDoRepository.UpdateToDo(ToDo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}

func (h *handlerToDo) DeleteToDo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	ToDo, err := h.ToDoRepository.GetOneToDo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ToDoRepository.DeleteToDo(ToDo, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}