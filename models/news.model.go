package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Id          int    `gorm:"type:int;unique" json:"id"`
	TitleDate   string `gorm:"type:varchar(99999);not null" json:"titleDate"`
	Description string `gorm:"type:varchar(99999);" json:"description"`
	Image       string `gorm:"type:varchar(99999);not null" json:"image"`
}
