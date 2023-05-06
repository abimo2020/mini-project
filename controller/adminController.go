package controller

import (
	"mini-project/models/payload"
	"mini-project/usecase"
	"net/http"

	"github.com/labstack/echo"
)

func DashboardAdminController(c echo.Context) error {
	dashboard := usecase.DashboardAdmin()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "success get admin dashboard",
		"dashboard": dashboard,
	})

}

func GetUsersController(c echo.Context) error {
	users, err := usecase.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetPetCategoryController(c echo.Context) error {
	category, err := usecase.GetPetCategory()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success get all category",
		"category": category,
	})
}

func GetPetsAdminController(c echo.Context) error {
	response, err := usecase.GetPets()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pets",
		"pets":    response,
	})
}

func CreatePetCategoryController(c echo.Context) error {
	var req payload.CreateCategory
	c.Bind(&req)
	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.CreatePetCategory(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create pet category",
	})
}
