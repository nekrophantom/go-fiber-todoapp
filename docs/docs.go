// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Reynold",
            "url": "https://github.com/nekrophantom",
            "email": "reynold@nekro.dev"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "description": "Get list of all todos",
                "produces": [
                    "application/json"
                ],
                "summary": "List Todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Todo"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Todo",
                "produces": [
                    "application/json"
                ],
                "summary": "Store Todo",
                "parameters": [
                    {
                        "description": "Task Name",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\n\"task\" : \"example\"\n}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get todo by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Show an Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an Todo By ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Update Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task Name",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string",
                            "example": "{\n\"task\" : \"example\"\n}"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Todo"
                        }
                    },
                    "400": {
                        "description": "Invalid Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an Todo By ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success Delete Todo",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.Todo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "is_done": {
                    "type": "boolean"
                },
                "task": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "fiber-todoapp.nekro.dev",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Todo App Go - Fiber",
	Description:      "Simple API using Go - Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
