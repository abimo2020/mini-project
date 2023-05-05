package controller

import (
	"mini-project/lib/database"
	"net/http"

	"github.com/labstack/echo"
)

func DashboardAdminController(c echo.Context) error {
	dashboard, err := database.DashboardAdmin(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success get admin dashboard",
		"dashboard": dashboard,
	})

}

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetPetCategoryController(c echo.Context) error {
	category, err := database.GetPetCategory(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all category",
		"category": category,
	})
}

func GetPetsAdminController(c echo.Context) error {
	pets, err := database.GetPetsAdmin(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pets",
		"pets":    pets,
	})
}

func CreatePetCategoryController(c echo.Context) error {
	err := database.CreatePetCategory(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create pet category",
	})
}
