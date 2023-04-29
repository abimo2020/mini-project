package route

import (
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

	e.POST("/login", controller.LoginController)
	e.POST("/logout", controller.LogoutController)

	e.POST("/register", controller.RegisterController)
	e.POST("/donate", controller.DonateController, m.JWTMiddlewareConfig)

	e.GET("/dashboard", controller.DashboardUserController, m.JWTMiddlewareConfig)

	pets := e.Group("/pets")
	pets.GET("", controller.GetPetsController)
	pets.GET("/:id", controller.GetPetController)
	pets.POST("/:id/adopt", controller.AdoptController, m.JWTMiddlewareConfig)

	p := e.Group("/profil")
	p.GET("", controller.GetProfilController, m.JWTMiddlewareConfig)
	p.PUT("/update-detail", controller.UpdateProfilDetailController, m.JWTMiddlewareConfig)
	p.PUT("/update", controller.UpdateProfilController, m.JWTMiddlewareConfig)
	p.DELETE("/delete", controller.DeleteUserController, m.JWTMiddlewareConfig)

	return e
}
