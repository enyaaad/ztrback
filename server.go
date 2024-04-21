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

type note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var notes = []note{
	{ID: "1", Title: "Buy food", Description: "Go shop and buy something"},
	{ID: "2", Title: "Do workout", Description: "Get six press cubes"},
	{ID: "3", Title: "Kill Kenny", Description: "He must die"},
}

func StartAPI() {

	app := fiber.New(config)

	// Create a new route with GET method
	app.Get("/app", func(c *fiber.Ctx) error {
		return c.JSON(notes)
	})

	app.Get("/projects", handlers.GetAllProjects)
	app.Get("/body", handlers.GetAllBackgrounds)
	app.Get("/socials", handlers.GetAllSocials)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
}
