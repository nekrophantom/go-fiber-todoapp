package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/go-fiber-todoapp/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my To do App API")
}

func main() {

	database.OpenConnection()

	app := fiber.New()

	app.Get("/", welcome)

	err := app.Listen("localhost:3000")

	if err != nil {
		panic(err)
	}
}