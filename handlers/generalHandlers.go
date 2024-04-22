package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllBackgrounds(c *fiber.Ctx) error {
	var backgrounds []models.Home
	if err := db.DB.Find(&backgrounds).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(backgrounds)
}

func GetRandomBackground(c *fiber.Ctx) error {
	var background models.Home

	if err := db.DB.Order("RANDOM()").Limit(1).First(&background).Error; err != nil {
		// Если запись не найдена, возвращаем ошибку
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		// Если произошла другая ошибка, возвращаем ошибку
		return err
	}
	return c.JSON(background)
}
