package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllNews(c *fiber.Ctx) error {
	var news []models.News
	if err := db.DB.Find(&news).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(news)
}

func GetNewsByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var news models.News
	if err := db.DB.First(&news, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("News not found")
		}

		return c.Status(fiber.StatusInternalServerError).JSON("Error while finding")

	}

	return c.JSON(news)
}
