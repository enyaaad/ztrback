package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Id          int    `gorm:"type:int;unique" json:"id"`
	Names       string `gorm:"type:varchar;not null" json:"title"`
	Preview     string `gorm:"type:varchar;not null" json:"preview"`
	Description string `gorm:"type:varchar;not null" json:"description"`
}
