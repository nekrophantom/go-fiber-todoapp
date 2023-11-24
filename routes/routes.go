package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/go-fiber-todoapp/controller"
	"gorm.io/gorm"
)

func SetupAPIRoutes(app *fiber.App, db *gorm.DB) {

	todoController := controller.NewTodoController(db)

	apiGroup := app.Group("/api")
	
	todosGroup := apiGroup.Group("/todos")
	todosGroup.Get("/", todoController.GetAllTodos)
	todosGroup.Get("/:id", todoController.GetTodoById)

}
