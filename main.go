package main

import (
	"fmt"

	"github.com/jackthepanda96/Belajar-Rest.git/config"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/user/data"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/user/delivery"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/user/usecase"
	"github.com/jackthepanda96/Belajar-Rest.git/infrastructure/database/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	e := echo.New()
	userData := data.New(db)
	useCase := usecase.New(userData)
	delivery.New(e, useCase)
	fmt.Println("Menjalankan program ....")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))

}
