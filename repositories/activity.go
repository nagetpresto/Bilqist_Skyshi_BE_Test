package repositories

import (
	"BE/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	GetAllActivity() ([]models.Activity, error)
	GetOneActivity(ID int) (models.Activity, error)
	CreateActivity(activity models.Activity) (models.Activity, error)
	UpdateActivity(activity models.Activity) (models.Activity, error)
	DeleteActivity(activity models.Activity, ID int) (models.Activity, error)
}

func RepositoryActivity(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllActivity() ([]models.Activity, error) {
	var activity []models.Activity
	err := r.db.Find(&activity).Error // add this code

	return activity, err
}

func (r *repository) GetOneActivity(ID int) (models.Activity, error) {
	var activity models.Activity
	err := r.db.First(&activity, ID).Error // add this code

	return activity, err
}

func (r *repository) CreateActivity(activity models.Activity) (models.Activity, error) {
	err := r.db.Create(&activity).Error

	return activity, err
}

func (r *repository) UpdateActivity(activity models.Activity) (models.Activity, error) {
	err := r.db.Save(&activity).Error

	return activity, err
}

func (r *repository) DeleteActivity(activity models.Activity, ID int) (models.Activity, error) {
	err := r.db.Delete(&activity, ID).Scan(&activity).Error

	return activity, err
}