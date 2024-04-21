package migrations

import (
	"austem/models"
	"austem/postreql/db"
	"fmt"
)

func init() {
	db.StartDB()
}

func Migrator() {
	err := db.DB.AutoMigrate(&models.Projects{}, &models.Home{}, &models.Socials{}, &models.Team{}, &models.Video{})
	if err != nil {
		return
	}
	fmt.Println("migrations complete")
}
