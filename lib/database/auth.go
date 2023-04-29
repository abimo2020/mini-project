package database

import (
	"mini-project/config"
	"mini-project/lib/cookie"
	"mini-project/middlewares"
	"mini-project/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	var err error
	user := models.User{}
	c.Bind(&user)

	if err = config.DB.Where("username = ?", user.Username).Where("password = ? ", user.Password).First(&user).Error; err != nil {
		if err = config.DB.Where("email = ?", user.Email).Where("password = ? ", user.Password).First(&user).Error; err != nil {
			return err
		}
	}
	token, err := middlewares.CreateToken(user.Username, user.Role, user.ID)

	cookie.CreateJWTCookies(c, token)

	if err != nil {
		return err
	}
	return nil
}

func Logout(c echo.Context) error {
	cookie, err := c.Cookie("JWTCookie")
	if err != nil {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "Not logged in yet")
	}
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.SetCookie(cookie)
	return nil
}

func Register(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func Authorization(c echo.Context) (string, uint) {
	cookie, _ := c.Cookie("JWTCookie")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	id := uint(claims["user_id"].(float64))
	return username, id
}
