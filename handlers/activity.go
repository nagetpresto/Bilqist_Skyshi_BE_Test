package handlers

import (
	dto "BE/dto/result"
	activitydto "BE/dto/activity"
	"BE/models"
	"BE/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerActivity struct {
	ActivityRepository repositories.ActivityRepository
}

func HandlerActivity(ActivityRepository repositories.ActivityRepository) *handlerActivity {
	return &handlerActivity{ActivityRepository}
}

func (h *handlerActivity) GetAllActivity(c echo.Context) error {
	activity, err := h.ActivityRepository.GetAllActivity()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: activity})
}

func (h *handlerActivity) GetOneActivity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	activity, err := h.ActivityRepository.GetOneActivity(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: activity})
}

func (h *handlerActivity) CreateActivity(c echo.Context) error {
	request := new(activitydto.ActivityRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	activity := models.Activity{
		Title:     	request.Title,
		Email:    	request.Email,
		CreatedAt: 	time.Now(),
	}

	data, err := h.ActivityRepository.CreateActivity(activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}

func (h *handlerActivity) UpdateActivity(c echo.Context) error {
	request := new(activitydto.UpdateActivityRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	activity, err := h.ActivityRepository.GetOneActivity(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Title != "" {
		activity.Title 		= request.Title
		activity.UpdateAt 	= time.Now()
	}

	data, err := h.ActivityRepository.UpdateActivity(activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}

func (h *handlerActivity) DeleteActivity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	activity, err := h.ActivityRepository.GetOneActivity(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ActivityRepository.DeleteActivity(activity, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Message: "Success", Data: data})
}