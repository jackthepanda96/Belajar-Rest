package delivery

import (
	"log"
	"net/http"

	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/common"
	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookUsecase domain.BookUseCase
}

func New(bu domain.BookUseCase) domain.BookHandler {
	return &bookHandler{
		bookUsecase: bu,
	}
}

func (bh *bookHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp BookInsertRequest
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := bh.bookUsecase.AddBook(common.ExtractData(c), tmp.ToDomain())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
		})

	}
}
