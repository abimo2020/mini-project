package controller

// dasboard, see profil, update profil, delete user,
import (
	"net/http"

	"mini-project/models/payload"
	"mini-project/usecase"

	"github.com/labstack/echo"
)

func DashboardUserController(c echo.Context) error {
	_, id := Authorization(c)
	dashboard := usecase.GetDashboardUser(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get the data",
		"dashboard": dashboard,
	})
}

func GetProfilController(c echo.Context) error {
	_, id := Authorization(c)
	user, err := usecase.GetProfil(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get the profil",
		"user":    user,
	})
}

func UpdateProfilDetailController(c echo.Context) error {
	var req payload.UpdateProfilDetail
	_, id := Authorization(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateProfilDetail(&req, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update profil detail",
	})
}

func UpdateProfilController(c echo.Context) error {
	var req payload.UpdateProfil
	_, id := Authorization(c)

	c.Bind(&req)

	if err := c.Validate(req); err != nil {
		return err
	}
	err := usecase.UpdateProfil(id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success update profil",
	})
}

func DeleteUserController(c echo.Context) error {
	_, id := Authorization(c)

	password := c.FormValue("password")
	err := usecase.DeleteUser(id, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
	})
}
