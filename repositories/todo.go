package repositories

import (
	"BE/models"
	"gorm.io/gorm"
)

type ToDoRepository interface {
	GetAllToDo(ActivityID int) ([]models.ToDo, error)
	GetOneToDo(ID int) (models.ToDo, error)
	CreateToDo(ToDo models.ToDo) (models.ToDo, error)
	UpdateToDo(ToDo models.ToDo) (models.ToDo, error)
	DeleteToDo(ToDo models.ToDo, ID int) (models.ToDo, error)
}

func RepositoryToDo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllToDo(ActivityID int) ([]models.ToDo, error) {
	var ToDo []models.ToDo
	err := r.db.Where("activity_id = ?", ActivityID).Find(&ToDo).Error // add this code
	return ToDo, err
}

func (r *repository) GetOneToDo(ID int) (models.ToDo, error) {
	var ToDo models.ToDo
	err := r.db.First(&ToDo, ID).Error // add this code

	return ToDo, err
}

func (r *repository) CreateToDo(ToDo models.ToDo) (models.ToDo, error) {
	err := r.db.Create(&ToDo).Error

	return ToDo, err
}

func (r *repository) UpdateToDo(ToDo models.ToDo) (models.ToDo, error) {
	err := r.db.Save(&ToDo).Error

	return ToDo, err
}

func (r *repository) DeleteToDo(ToDo models.ToDo, ID int) (models.ToDo, error) {
	err := r.db.Delete(&ToDo, ID).Scan(&ToDo).Error

	return ToDo, err
}