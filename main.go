package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Nama     string
	Email    string
	Password string
}

var (
	arrData []User
)

func GetAll(c echo.Context) error {
	res := map[string]interface{}{
		"message": "Get all data",
		"data":    arrData,
	}
	return c.JSON(http.StatusOK, res)
}

func InsertUser(c echo.Context) error {
	var tmp User
	err := c.Bind(&tmp)
	if err != nil {
		log.Println("Cannot parse input to object", err.Error())
		return c.JSON(http.StatusInternalServerError, "Error dari server")
	}

	arrData = append(arrData, tmp)
	res := map[string]interface{}{
		"message": "Success input data",
		"data":    tmp,
	}
	return c.JSON(http.StatusOK, res)
}

func GetSpecificUser(c echo.Context) error {
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}

	res := map[string]interface{}{
		"message": "Get all data",
		"data":    arrData[cnv-1],
	}
	return c.JSON(http.StatusOK, res)

}

func main() {
	e := echo.New()
	e.GET("/user", GetAll)
	e.POST("/user", InsertUser)
	e.GET("/user/:id", GetSpecificUser)

	fmt.Println("Menjalankan program ....")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
