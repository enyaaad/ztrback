package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetVideo(c *fiber.Ctx) error {
	title := c.Params("id")
	season := c.Params("season")
	series := c.Params("series_number")

	var video models.Video
	if err := db.DB.Table("videos").Where("title = ? AND season = ? AND series_number = ?", title, season, series).First(&video).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("Project not found")
		}

		return c.Status(fiber.StatusInternalServerError).JSON("Error while finding")

	}

	return c.JSON(video)
}
