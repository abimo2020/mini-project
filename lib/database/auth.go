package database

import (
	"mini-project/config"
	"mini-project/lib/cookie"
	"mini-project/middlewares"
	"mini-project/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func Login(c echo.Context) (interface{}, error) {
	var err error
	user := models.User{}
	c.Bind(&user)

	if err = config.DB.Where("username = ?", user.Username).Where("password = ? ", user.Password).First(&user).Error; err != nil {
		if err = config.DB.Where("email = ?", user.Email).Where("password = ? ", user.Password).First(&user).Error; err != nil {
			return nil, err
		}
	}
	token, err := middlewares.CreateToken(user.Username, user.Role, user.ID)

	cookie.CreateJWTCookies(c, token)

	if err != nil {
		return nil, err
	}
	return token, nil
}

func Register(c echo.Context) (interface{}, error) {
	user := models.User{}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func Authorization(c echo.Context) (string, uint) {
	cookie, _ := c.Cookie("JWTCookie")
	token, _ := jwt.Parse(cookie.Value, nil)
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	id := uint(claims["user_id"].(float64))
	return username, id
}
