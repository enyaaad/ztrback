package models

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Id          int    `gorm:"type:int;unique" json:"id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	InFuture    bool   `gorm:"type:bool;not null" json:"in_future"`
}
