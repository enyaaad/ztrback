package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllSocials(c *fiber.Ctx) error {
	var socials []models.Socials
	if err := db.DB.Find(&socials).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(socials)
}
