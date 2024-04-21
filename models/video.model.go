package models

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id           int    `gorm:"type:int;unique" json:"id"`
	Title        string `gorm:"type:varchar(255);not null" json:"title"`
	Season       int    `gorm:"type:integer;not null" json:"season"`
	SeriesNumber int    `gorm:"type:integer;not null" json:"series_number"`
	SeriesURL    string `gorm:"type:varchar(255);not null" json:"series_url"`
}
