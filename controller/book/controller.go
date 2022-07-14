package book

import (
	"log"
	"net/http"

	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	Model model.BookModel
}

func (bc *BookController) InsertBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.Book

		err := c.Bind(&input)
		if err != nil {
			log.Println("Kesalahan input dari user")
			return c.JSON(http.StatusBadRequest, "Incorrect input from user")
		}

		data := bc.Model.Insert(input)

		if data.ID == 0 {
			log.Println("gagal input")
			return c.JSON(http.StatusInternalServerError, "Cannot insert to database")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success add book",
			"data":    data,
		})
	}
}

func (bc *BookController) GetAllBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		data := bc.Model.GetAll()

		if data == nil {
			log.Println("Terdapat error saat mengambil data")
			return c.JSON(http.StatusInternalServerError, "problem from database")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success get all book",
			"data":    data,
		})
	}
}
