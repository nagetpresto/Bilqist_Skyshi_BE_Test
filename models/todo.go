package models

import (
	"time"
)

type ToDo struct {
	ID       		int    		`json:"id" gorm:"primary_key:auto_increment"`
	Title     		string 		`json:"title" gorm:"type: varchar(255)"`
	ActivityID     	int         `json:"activity_group_id"`
    Activity       	Activity  	`json:"-"`
	IsActive    	bool 		`json:"is_active" gorm:"type: boolean"`
	Priority     	string 		`json:"priority" gorm:"type: varchar(255)"`
	CreatedAt 		time.Time 	`gorm:"type:datetime" json:"created_at"`
	UpdateAt 		time.Time 	`gorm:"type:datetime" json:"updated_at"`
}