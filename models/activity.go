package models

import (
	"time"
)

type Activity struct {
	ID       		int    		`json:"id" gorm:"primary_key:auto_increment"`
	Title     		string 		`json:"title" gorm:"type: varchar(255)"`
	Email    		string 		`json:"email" gorm:"type: varchar(255)"`
	CreatedAt 		time.Time 	`gorm:"type:datetime" json:"created_at"`
	UpdateAt 		time.Time 	`gorm:"type:datetime" json:"updated_at"`
}