package main

import (
	"fmt"
	"net/http"
	"todolist/model"
	"todolist/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//se crean las http requests
func getAllTareas(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		ToDos := []*model.ToDos{}
		result := db.Order("id asc").Find(&ToDos)

		if result.Error != nil {
			fmt.Print(result.Error)
		}
		return c.JSON(http.StatusOK, ToDos)
	}
}

func getOneTarea(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var todo model.ToDos

		result := db.First(&todo, id)
		if result.Error != nil {
			fmt.Print(result.Error)
		}
		return c.JSON(http.StatusOK, todo)
	}
}

func postTarea(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		tarea := c.FormValue("tarea")
		todo := &model.ToDos{Tarea: tarea}

		result := db.Create(&todo)
		if result.Error != nil {
			fmt.Print(result.Error)
		}
		return c.JSON(http.StatusOK, todo)
	}
}

func putTarea(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		tarea := c.FormValue("tarea")
		var tareas model.ToDos
		resultGet := db.First(&tarea, id)

		if resultGet.Error != nil {
			fmt.Print(resultGet.Error)
		}

		tareas.Tarea = tarea

		if resultGet.Error != nil {
			fmt.Print(resultGet.Error)
		}
		return c.JSON(http.StatusOK, tarea)
	}
}

func deleteTarea(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		result := db.Delete(&model.ToDos{}, id)

		if result.Error != nil {
			fmt.Print(result.Error)
		}
		return c.String(http.StatusOK, "tarea eliminada")
	}
}

func main() {
	db := utils.InitConnection()

	e := echo.New()

	e.GET("/ToDos", getAllTareas(db))

	e.GET("ToDos/:id", getOneTarea(db))

	e.POST("/ToDos", postTarea(db))

	e.PUT("ToDos/:id", putTarea(db))

	e.DELETE("/ToDos/:id", deleteTarea(db))
	fmt.Println("iniciando el servidor")
	e.Logger.Fatal(e.Start(":3306"))

}
