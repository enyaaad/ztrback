package server

import (
	"austem/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	log "github.com/sirupsen/logrus"
)

var config = fiber.Config{
	ServerHeader: "My Server", // add custom server header
}

func StartAPI() {

	app := fiber.New(config)

	app.Use(func(c *fiber.Ctx) error {
		// Установка заголовков для разрешения запросов от любых источников
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Accept-Encoding", "gzip")

		// Пропуск запросов метода OPTIONS, так как он используется для предварительной проверки CORS
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusNoContent)
		}

		// Продолжаем выполнение запроса
		return c.Next()
	})

	v1 := app.Group("/api")

	v1.Get("/projects", handlers.GetAllProjects)
	v1.Get("/projects/:id", handlers.GetProjectByID)
	v1.Get("/projects/series/:id", handlers.GetProjectByID)

	v1.Get("/body", handlers.GetAllBackgrounds)
	v1.Get("/body/random", handlers.GetRandomBackground)

	v1.Get("/socials", handlers.GetAllSocials)

	v1.Get("/video/:id/:season/:series_number", handlers.GetVideo)

	v1.Get("/news", handlers.GetAllNews)
	v1.Get("/news:id", handlers.GetNewsByID)

	app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
}
