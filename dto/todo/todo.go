package tododto

import (
	"time"
	"BE/models"
)

type ToDoRequest struct {
	Title     		string 		`json:"title" gorm:"type: varchar(255)" validate:"required"`
	ActivityID     	int         `json:"activity_group_id" validate:"required"`
    IsActive    	bool 		`json:"is_active" gorm:"type: boolean" validate:"required"`
	Priority     	string 		`json:"priority" gorm:"type: varchar(255)" validate:"required"`
	CreatedAt 		time.Time 	`gorm:"type:datetime" json:"created_at"`
	UpdateAt 		time.Time 	`gorm:"type:datetime" json:"updated_at"`
}

type UpdateToDoRequest struct {
	Title     		string 		`json:"title" gorm:"type: varchar(255)"`
	IsActive    	bool 		`json:"is_active" gorm:"type: boolean"`
	Priority     	string 		`json:"priority" gorm:"type: varchar(255)"`
}

type ToDoResponse struct {
	ID       		int    		`json:"id" gorm:"primary_key:auto_increment"`
	Title     		string 		`json:"title" gorm:"type: varchar(255)"`
	ActivityID     	int         `json:"activity_group_id"`
    Activity       	models.Activity  	`json:"-"`
	IsActive    	bool 		`json:"is_active" gorm:"type: boolean"`
	Priority     	string 		`json:"priority" gorm:"type: varchar(255)"`
	CreatedAt 		time.Time 	`gorm:"type:datetime" json:"created_at"`
	UpdateAt 		time.Time 	`gorm:"type:datetime" json:"updated_at"`
}