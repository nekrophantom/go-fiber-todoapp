package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nekrophantom/go-fiber-todoapp/models"
	"gorm.io/gorm"
)

type TodoController struct {
	DB *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{
		DB: db,
	}
}

// Get All
// @Summary List Todos
// @Description Get list of all todos
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func (c *TodoController) GetAllTodos(ctx *fiber.Ctx) error {
	var todos []models.Todo
	c.DB.Find(&todos)
	return ctx.JSON(todos)
}

// Get by Id
// @Summary Show an Todo
// @Description Get todo by ID
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 404  {string}  string "Todo not found"
// @Router /todos/{id} [get]
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
// @Summary Store Todo
// @Description Create a Todo
// @Produce json
// @Param task body string true "Task Name" SchemaExample({\n"task" : "example"\n})
// @Success 200 {object} models.Todo
// @Failure 400  {string}  string "Invalid Request"
// @Router /todos [post]
func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
	var todo models.Todo

	if err := ctx.BodyParser(&todo); err != nil {
		return ctx.Status(400).SendString("Invalid Request")
	}

	c.DB.Create(&todo)
	return ctx.JSON(todo)
}

// Update Todo by Id
// @Summary Update Todo
// @Description Update an Todo By ID
// @Produce json
// @Param id path int true "Todo ID"
// @Param task body string true "Task Name" SchemaExample({\n"task" : "example"\n})
// @Success 200 {object} models.Todo
// @Failure 400  {string}  string "Invalid Request"
// @Failure 404  {string}  string "Todo not found"
// @Router /todos/{id} [put]
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
// @Summary Delete Todo
// @Description Delete an Todo By ID
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204 {string} string "Success Delete Todo" 
// @Failure 404  {string}  string "Todo not found"
// @Router /todos/{id} [delete]
func (c *TodoController) DeleteTodo(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var todo models.Todo
	result := c.DB.First(&todo, id)

	if result.Error != nil {
		return ctx.Status(404).SendString("Todo not found")
	}

	c.DB.Delete(&todo)
	return ctx.Status(200).SendString("Success Delete Todo")
}
