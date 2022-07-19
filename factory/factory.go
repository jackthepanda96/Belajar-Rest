package factory

import (
	"gorm.io/gorm"

	"github.com/go-playground/validator/v10"
	bd "github.com/jackthepanda96/Belajar-Rest.git/feature/book/data"
	bookDelivery "github.com/jackthepanda96/Belajar-Rest.git/feature/book/delivery"
	"github.com/labstack/echo/v4"

	bs "github.com/jackthepanda96/Belajar-Rest.git/feature/book/usecase"

	ud "github.com/jackthepanda96/Belajar-Rest.git/feature/user/data"

	userDelivery "github.com/jackthepanda96/Belajar-Rest.git/feature/user/delivery"

	us "github.com/jackthepanda96/Belajar-Rest.git/feature/user/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	bookData := bd.New(db)
	bookCase := bs.New(bookData)
	bookHandler := bookDelivery.New(bookCase)
	bookDelivery.RouteBook(e, bookHandler)
}
