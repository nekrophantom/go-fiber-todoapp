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

// Get by Id
func (c *TodoController) GetTodoById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var todo models.Todo
	result := c.DB.First(&todo, id)

	if result.Error != nil {
		return ctx.Status(404).SendString("Todo not found")
	}

	return ctx.JSON(todo)
}

// Create a new Todo
func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo models.Todo
	
	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).SendString("Invalid Request")
	}

	c.DB.Create(&todo)
	return ctx.JSON(todo)
}

// Update Todo by Id
func (c *TodoController) UpdateTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var todo models.Todo
	result := c.DB.First(&todo, id)

	if result.Error != nil {
		return ctx.Status(404).SendString("Todo not found")
	}

	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).SendString("Invalid Request")
	}

	c.DB.Save(&todo)
	return ctx.JSON(todo)
}

// Delete Todo by Id
func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var todo models.Todo
	result := c.DB.First(&todo, id)

	if result.Error != nil {
		return ctx.Status(404).SendString("Todo not found")
	}

	c.DB.Delete(&todo)
	return ctx.SendStatus(204)
}