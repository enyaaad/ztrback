package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllBackgrounds(c *fiber.Ctx) error {
	var backgrounds []models.Home
	if err := db.DB.Find(&backgrounds).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(backgrounds)
}
