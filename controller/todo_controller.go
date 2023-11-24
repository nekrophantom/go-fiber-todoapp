package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/go-fiber-todoapp/models"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController{
	return &TodoController{
		DB: db,
	}
}

// Get All
func (c *TodoController) GetAllTodos(ctx *fiber.Ctx) error {
	var todos []models.Todo
	c.DB.Find(&todos)
	return ctx.JSON(todos)
}