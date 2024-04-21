package models

import (
	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	Id           int    `gorm:"type:int;unique" json:"id"`
	Title        string `gorm:"type:varchar(255);not null" json:"title"`
	Image        string `gorm:"type:varchar(255);not null" json:"image"`
	Description  string `gorm:"type:varchar(255);not null" json:"description"`
	VideoPreview string `gorm:"type:varchar;not null" json:"video_preview"`
	InFuture     bool   `gorm:"type:bool;not null" json:"in_future"`
}
