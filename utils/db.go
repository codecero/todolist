package utils

import (
	"log"
	"todolist/config"
	"todolist/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection() *gorm.DB {
	dsn := config.GetDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error de coneccion!!")
	}
	db.AutoMigrate(&model.ToDos{})
	ToDo := &model.ToDos{Tarea: "Comprar pan", Completada: false}

	db.Create(&ToDo)

	return db
}
