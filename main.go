package main

import (
	"github.com/NaufalFarros/golang_fiber_restapi/models"
	"github.com/gofiber/fiber/v2"
)

func main() {

	models.SetupDatabase()

	app := fiber.New()

	app.Get("/api/books", controllers.GetBooks)
	app.Get("/api/books/:id", controllers.Show)
	app.Post("/api/books", controllers.Create)
	app.Put("/api/books/:id", controllers.Update)
	app.Delete("/api/books/:id", controllers.Delete)
}
