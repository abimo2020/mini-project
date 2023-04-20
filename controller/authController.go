package controller

import (
	"mini-project/lib/database"
	"mini-project/models"
	"net/http"

	"github.com/labstack/echo"
)

func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	token, err := database.LoginUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login user",
		"token":  token,
	})
}

func RegisterUserController(c echo.Context) error {
	user, err := database.Register(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}
