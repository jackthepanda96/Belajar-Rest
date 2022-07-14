package main

import (
	"fmt"

	"github.com/jackthepanda96/Belajar-Rest.git/controller/book"
	"github.com/jackthepanda96/Belajar-Rest.git/controller/user"
	"github.com/jackthepanda96/Belajar-Rest.git/database/mysql"
	"github.com/jackthepanda96/Belajar-Rest.git/middlewares"
	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := mysql.InitDB()
	mysql.MigrateData(db)
	e := echo.New()
	loginModel := model.LoginModel{DB: db}
	userModel := model.UserModel{DB: db}
	userController := user.UserController{Model: userModel, Auth: loginModel}

	bookModel := model.BookModel{DB: db}
	bookController := book.BookController{Model: bookModel}

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS()) //WAJIB!!
	// e.Use(middleware.Logger()) //WAJIB!!
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// e.Use(middleware.BasicAuth(middlewares.CekAuth))

	e.POST("/login", userController.Login())

	user := e.Group("/user")
	user.GET("", userController.GetAll(), middleware.JWT([]byte("R4h@s1A!")))
	user.POST("", userController.InsertUser())
	user.GET("/:id", userController.GetSpecificUser())
	user.PUT("/:id", userController.UpdateUser())
	user.DELETE("/:id", userController.DeleteUser())

	book := e.Group("/book", middleware.BasicAuth(middlewares.CekAuth))
	book.GET("", bookController.GetAllBook())
	book.POST("", bookController.InsertBook())

	fmt.Println("Menjalankan program ....")
	e.Logger.Fatal(e.Start(":8000"))

}
