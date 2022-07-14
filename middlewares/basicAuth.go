package middlewares

import "github.com/labstack/echo/v4"

func CekAuth(s1, s2 string, ctx echo.Context) (bool, error) {
	if s1 == "Admin" && s2 == "Admin" {
		return true, nil
	}

	return false, nil
}
