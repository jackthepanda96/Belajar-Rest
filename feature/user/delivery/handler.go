package delivery

import (
	"log"
	"net/http"
	"strings"

	"github.com/jackthepanda96/Belajar-Rest.git/config"
	"github.com/jackthepanda96/Belajar-Rest.git/domain"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/book/delivery/middlewares"
	"github.com/jackthepanda96/Belajar-Rest.git/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

func New(e *echo.Echo, us domain.UserUseCase) {
	handler := &userHandler{
		userUsecase: us,
	}
	e.POST("/users", handler.InsertUser())
	e.GET("/users", handler.GetAllUser())
	e.GET("/profile", handler.GetProfile(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
			"token":   common.GenerateToken(data.ID),
		})
	}
}

func (uh *userHandler) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := uh.userUsecase.GetAll()

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("User Handler", err)
				c.JSON(http.StatusNotFound, err.Error())
			} else if strings.Contains(err.Error(), "retrieve") {
				log.Println("User Handler", err)
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all user data",
			"data":    data,
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
		})
	}
}
