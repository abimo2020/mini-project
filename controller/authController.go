package controller

import (
	"mini-project/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func LoginController(c echo.Context) error {

	err := database.Login(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login user",
	})
}

func RegisterController(c echo.Context) error {
	err := database.Register(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to register",
	})
}

func LogoutController(c echo.Context) error {
	err := database.Logout(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to logout",
	})
}
