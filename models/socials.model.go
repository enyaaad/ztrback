package models

import (
	"gorm.io/gorm"
)

type Socials struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:varchar(255);not null" json:"description"`
	Url         string `gorm:"type:varchar(255);not null" json:"url"`
	LogoUrl     string `gorm:"type:varchar(255);not null" json:"logo_url"`
	ImageUrl    string `gorm:"type:varchar(255);not null" json:"image_url"`
}
