package main

import (
	"fmt"
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/controller/user"
	"github.com/jackthepanda96/Belajar-Rest.git/database/mysql"
	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitDB()
	e := echo.New()
	model := model.UserModel{DB: db}
	controller := user.UserController{Model: model}

	e.GET("/user", controller.GetAll())
	e.POST("/user", controller.InsertUser())
	e.GET("/user/:id", controller.GetSpecificUser())
	e.PUT("/user/:id", controller.UpdateUser())
	e.DELETE("/user/:id", controller.DeleteUser())

	fmt.Println("Menjalankan program ....")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
