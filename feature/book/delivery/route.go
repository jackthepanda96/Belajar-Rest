package delivery

import (
	"github.com/jackthepanda96/Belajar-Rest.git/config"
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/book/delivery/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteBook(e *echo.Echo, bc domain.BookHandler) {
	e.POST("/book", bc.InsertUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
