package handlers

import (
	"austem/models"
	"austem/postreql/db"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllProjects(c *fiber.Ctx) error {
	var projects []models.Projects

	if err := db.DB.Model(&projects).Preload("Seasons.Videos").Find(&projects).Error; err != nil {
		// Обработка ошибки
		return nil
	}

	return c.JSON(projects)
}

func GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var project models.Projects
	if err := db.DB.Model(&project).Preload("Seasons.Videos").First(&project, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON("Project not found")
		}

		return c.Status(fiber.StatusInternalServerError).JSON("Error while finding")

	}

	return c.JSON(project)
}
