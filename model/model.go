package model

import "gorm.io/gorm"

type ToDos struct {
	gorm.Model
	Tarea      string
	Completada bool
}
