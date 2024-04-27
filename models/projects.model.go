package models

type Projects struct {
	ID           uint     `gorm:"primaryKey" json:"id"`
	Title        string   `gorm:"type:varchar;not null" json:"title"`
	Image        string   `gorm:"type:varchar;not null" json:"image"`
	Description  string   `gorm:"type:varchar;not null" json:"description"`
	VideoPreview string   `gorm:"type:varchar;not null" json:"video_preview"`
	Seasons      []Season `gorm:"foreignKey:ProjectsID" json:"seasons"`
	InFuture     bool     `gorm:"type:bool;not null" json:"in_future"`
}
