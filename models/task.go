package models

import "gorm.io/gorm"

// Definimos las constantes (el "diccionario" de estados)
const (
	StatusTodo       = "todo"
	StatusInProgress = "in-progress"
	StatusDone       = "done"
)

type Task struct {
	gorm.Model
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:todo"`
}