package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Task      string `json:"task"`
	IsDone    bool   `json:"is_done"`
}
