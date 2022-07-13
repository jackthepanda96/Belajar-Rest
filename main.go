package main

import (
	"fmt"
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/controller/book"
	"github.com/jackthepanda96/Belajar-Rest.git/controller/user"
	"github.com/jackthepanda96/Belajar-Rest.git/database/mysql"
	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitDB()
	mysql.MigrateData(db)
	e := echo.New()
	userModel := model.UserModel{DB: db}
	userController := user.UserController{Model: userModel}

	bookModel := model.BookModel{DB: db}
	bookController := book.BookController{Model: bookModel}

	user := e.Group("/user")
	user.GET("", userController.GetAll())
	user.POST("", userController.InsertUser())
	user.GET("/:id", userController.GetSpecificUser())
	user.PUT("/:id", userController.UpdateUser())
	user.DELETE("/:id", userController.DeleteUser())

	book := e.Group("/book")
	book.GET("", bookController.GetAllBook())
	book.POST("", bookController.InsertBook())

	fmt.Println("Menjalankan program ....")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
