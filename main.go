package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	arrData []User
)

func cariData(id int) (int, bool) {
	for idx, val := range arrData {
		if val.ID == id {
			return idx, true
		}
	}

	return -1, false
}

func GetAll(c echo.Context) error {
	res := map[string]interface{}{
		"message": "Get all data",
		"data":    arrData,
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

	if cnv > len(arrData) {
		log.Println("Index out of range")
		return c.JSON(http.StatusInternalServerError, "Index out of range")
	}

	idxData, isFound := cariData(cnv)

	if !isFound {
		log.Println("Data not found")
		return c.JSON(http.StatusNotFound, "Data not found")
	}

	res := map[string]interface{}{
		"message": "Get all data",
		"data":    arrData[idxData],
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

	if len(arrData) == 0 {
		tmp.ID = 1
	} else {
		tmp.ID = arrData[len(arrData)-1].ID + 1
	}

	arrData = append(arrData, tmp)
	res := map[string]interface{}{
		"message": "Success input data",
		"data":    tmp,
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	param := c.Param("id")
	cnv, err := strconv.Atoi(param)
	if err != nil {
		log.Println("Cannot convert to int", err.Error())
		return c.JSON(http.StatusInternalServerError, "cannot convert id")
	}

	if cnv > len(arrData) {
		log.Println("Index out of range")
		return c.JSON(http.StatusInternalServerError, "Index out of range")
	}

	idxData, isFound := cariData(cnv)

	if !isFound {
		log.Println("Data not found")
		return c.JSON(http.StatusNotFound, "Data not found")
	}

	left := arrData[:idxData]
	right := arrData[idxData+1:]
	arrData = left
	arrData = append(arrData, right...)

	res := map[string]interface{}{
		"message": "Success delete data",
	}
	return c.JSON(http.StatusOK, res)
}

func main() {
	e := echo.New()
	e.GET("/user", GetAll)
	e.POST("/user", InsertUser)
	e.GET("/user/:id", GetSpecificUser)
	e.DELETE("/user/:id", DeleteUser)

	fmt.Println("Menjalankan program ....")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
