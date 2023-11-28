package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nekrophantom/go-fiber-todoapp/database"
	"github.com/nekrophantom/go-fiber-todoapp/routes"

	"github.com/gofiber/swagger"
	_ "github.com/nekrophantom/go-fiber-todoapp/docs"
)

// @title Todo App Go - Fiber
// @version 1
// @Description Simple API using Go - Fiber

// @contact.name Reynold
// @contact.url https://github.com/nekrophantom
// @contact.email reynold@nekro.dev

// @host fiber-todoapp.nekro.dev
// @BasePath /api

func main() {

	database.OpenConnection()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetupAPIRoutes(app, database.DB)
	
	app.Get("/*", swagger.HandlerDefault) // default

	err := app.Listen("localhost:3000")

	if err != nil {
		panic(err)
	}
}