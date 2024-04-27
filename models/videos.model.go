package models

type Season struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	ProjectsID uint    `json:"ProjectsID"`
	Number     int     // Номер сезона
	Videos     []Video `gorm:"foreignKey:SeasonID" json:"videos"` // Связь один ко многим с эпизодами
}

type Video struct {
	ID           uint   `gorm:"primaryKey"`
	SeasonID     uint   `json:"SeasonID"`
	SeriesNumber int    `gorm:"type:integer;not null;" json:"series_number"`
	SeriesURL    string `gorm:"type:varchar(255);not null" json:"series_url"`
}
