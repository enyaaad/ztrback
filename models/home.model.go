package models

import (
	"gorm.io/gorm"
)

type Home struct {
	gorm.Model
	BackgroundURL string `gorm:"type:varchar(255);not null" json:"backgroundURL"`
}
