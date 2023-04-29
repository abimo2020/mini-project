package database

import (
	"mini-project/config"
	"mini-project/models"

	"github.com/labstack/echo"
)

// func DashboardAdmin(c echo.Context) (interface{}, error) {

// }

func CreatePetCategory(c echo.Context) (interface{}, error) {
	var category models.PetCategory
	c.Bind(&category)
	if err := config.DB.Save(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}
