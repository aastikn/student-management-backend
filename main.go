package main

import (
	"student-management/database"
	"student-management/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/student/:id",routes.GetStudent)
	app.Post("/student",routes.CreateStudent)
}

func main() {
	app := fiber.New()

	database.ConnectToDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	SetupRoutes(app)

	app.Listen(":3000")

}