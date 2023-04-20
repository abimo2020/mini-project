package route

import (
	"mini-project/constants"
	"mini-project/controller"
	m "mini-project/middlewares"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	m.LoggerMiddleware(e)
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/login", controller.LoginUserController)
	e.POST("/register", controller.RegisterUserController)

	u := e.Group("/users")
	u.GET("", controller.GetUsersController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.GET("/:id", controller.GetUserController, middleware.JWT([]byte(constants.SECRET_JWT)))

	u.DELETE("/:id", controller.DeleteUserController, middleware.JWT([]byte(constants.SECRET_JWT)))
	u.PUT("/:id", controller.UpdateUserController, middleware.JWT([]byte(constants.SECRET_JWT)))

	return e
}
