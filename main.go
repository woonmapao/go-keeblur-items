package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/woonmapao/go-keeblur-items/controllers"
)

func main() {
	if err := controllers.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	items := app.Group("/items")

	items.Get("/all", controllers.GetAll)

	items.Get("/type/:type", controllers.GetByType)

	items.Get("/id/:id", controllers.GetByID)

	log.Fatal(app.Listen(":3000"))
}
