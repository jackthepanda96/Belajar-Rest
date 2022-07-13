package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func (uc *UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmp := uc.Model.GetAll()

		if tmp == nil {
			return c.JSON(http.StatusInternalServerError, "error from database")
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    tmp,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) GetSpecificUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		data := uc.Model.GetSpecific(cnv)

		if data.ID == 0 {
			return c.JSON(http.StatusBadRequest, "no data")
		}

		res := map[string]interface{}{
			"message": "Get all data",
			"data":    data,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp model.User
		err := c.Bind(&tmp)
		if err != nil {
			log.Println("Cannot parse input to object", err.Error())
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		data := uc.Model.Insert(tmp)

		if data.ID == 0 {
			return c.JSON(http.StatusInternalServerError, "Error dari server")
		}

		res := map[string]interface{}{
			"message": "Success input data",
			"data":    data,
		}
		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		qry := map[string]interface{}{}
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		var tmp model.User
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
		data := uc.Model.Update(cnv, tmp)

		if data.ID == 0 {
			return c.JSON(http.StatusInternalServerError, "cannot update")
		}

		res := map[string]interface{}{
			"message": "Success update data",
			"data":    data,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (uc *UserController) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		cnv, err := strconv.Atoi(param)
		if err != nil {
			log.Println("Cannot convert to int", err.Error())
			return c.JSON(http.StatusInternalServerError, "cannot convert id")
		}

		if !uc.Model.Delete(cnv) {
			return c.JSON(http.StatusInternalServerError, "cannot delete")
		}

		res := map[string]interface{}{
			"message": "Success delete data",
		}
		return c.JSON(http.StatusOK, res)
	}
}
