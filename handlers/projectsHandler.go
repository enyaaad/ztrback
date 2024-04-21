package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"github.com/gofiber/fiber/v2"
)

func GetAllProjects(c *fiber.Ctx) error {
	var projects []models.Projects
	if err := db.DB.Find(&projects).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(projects)
}
