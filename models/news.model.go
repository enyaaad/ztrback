package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Id          int    `gorm:"type:int;unique" json:"id"`
	TitleDate   string `gorm:"type:varchar(255);not null" json:"titleDate"`
	Description string `gorm:"type:varchar(255);" json:"description"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
}
