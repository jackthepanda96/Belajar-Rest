package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" form:"id"`
	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	dataNumber int
	// DB         *gorm.DB
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", "root", "", "localhost", 3307, "echorm")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	db.AutoMigrate(User{})
	return db
}

func GetAll(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp []User
		err := db.Find(&tmp).Error

		if err != nil {
			log.Println("Cannot retrive object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func GetSpecificUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp User

		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		err = db.Where("ID = ?", cnv).First(&tmp).Error
		if err != nil {
			log.Println("There is a problem with data", err.Error())
			return c.JSON(http.StatusBadRequest, "no data")
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func InsertUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp User
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		if dataNumber == 0 {
			tmp.ID = 1
		} else {
			dataNumber++
			tmp.ID = dataNumber
		}

		err = db.Create(&tmp).Error
		if err != nil {
			log.Println("Cannot create object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		res := map[string]interface{}{
			"message": "Success input data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func UpdateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		qry := map[string]interface{}{}
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp User
		err = c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		if tmp.Nama != "" {
			qry["nama"] = tmp.Nama
		}

		if tmp.Email != "" {
			qry["email"] = tmp.Email
		}

		if tmp.Password != "" {
			qry["password"] = tmp.Password
		}
		var ret User
		err = db.Model(&ret).Where("ID = ?", cnv).Updates(qry).Error
		if err != nil {
			log.Println("Cannot update data", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot update")
		}

		res := map[string]interface{}{
			"message": "Success update data",
			"data":    ret,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func DeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		err = db.Where("ID = ?", cnv).Delete(&User{}).Error
		if err != nil {
			log.Println("Cannot delete data", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		res := map[string]interface{}{
			"message": "Success delete data",
		}
		return c.JSON(http.StatusOK, res)
	}
}

func main() {
	db := InitDB()
	e := echo.New()
	e.GET("/user", GetAll(db))
	e.POST("/user", InsertUser(db))
	e.GET("/user/:id", GetSpecificUser(db))
	e.PUT("/user/:id", UpdateUser(db))
	e.DELETE("/user/:id", DeleteUser(db))

	fmt.Println("Menjalankan program ....")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
