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
	e.POST("/login-admin", controller.LoginAdminController)

	e.POST("/register", controller.RegisterUserController)

	u := e.Group("/users")
	u.GET("", controller.GetUsersController, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(constants.SECRET_KEY),
		TokenLookup:   "cookie:JWTCookie",
		AuthScheme:    "user",
	}))
	u.GET("/:id", controller.GetUserController, middleware.JWT([]byte(constants.SECRET_KEY)))
	u.DELETE("/:id", controller.DeleteUserController, middleware.JWT([]byte(constants.SECRET_KEY)))
	u.PUT("/:id", controller.UpdateUserController, middleware.JWT([]byte(constants.SECRET_KEY)))

	a := e.Group("/admin")
	a.GET("", controller.GetUsersController, middleware.JWT([]byte(constants.SECRET_KEY)))
	a.GET("/:id", controller.GetUserController, middleware.JWT([]byte(constants.SECRET_KEY)))
	a.DELETE("/:id", controller.DeleteUserController, middleware.JWT([]byte(constants.SECRET_KEY)))
	a.PUT("/:id", controller.UpdateUserController, middleware.JWT([]byte(constants.SECRET_KEY)))

	p := e.Group("/pets")
	p.GET("", controller.GetPetsController)
	p.GET("/:id", controller.GetPetController)
	p.POST("", controller.CreatePetController)
	p.DELETE("/:id", controller.DeletePetController)

	return e
}
