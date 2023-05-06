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

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)
	e.POST("/logout", controller.LogoutController)

	e.GET("/dashboard", controller.DashboardUserController, m.JWTMiddlewareConfig, m.IsUser)
	e.POST("/donate", controller.DonateController, m.JWTMiddlewareConfig, m.IsUser)

	e.GET("/my-donates", controller.GetDonateListController, m.JWTMiddlewareConfig, m.IsUser)
	e.PUT("/my-donates/:id/status", controller.UpdatePetStatusController, m.JWTMiddlewareConfig, m.IsUser)
	e.PUT("/my-donates/:id", controller.UpdatePetController, m.JWTMiddlewareConfig, m.IsUser)
	e.DELETE("/my-donates/:id", controller.DeletePetController, m.JWTMiddlewareConfig, m.IsUser)

	e.GET("/my-adopts", controller.GetAdoptListController, m.JWTMiddlewareConfig, m.IsUser)

	pets := e.Group("/pets")
	pets.GET("", controller.GetPetsController)
	pets.GET("/:id", controller.GetPetController)
	pets.POST("/:id/adopt", controller.AdoptController, m.JWTMiddlewareConfig)

	p := e.Group("/profil")
	p.Use(m.JWTMiddlewareConfig, m.IsUser)
	p.GET("", controller.GetProfilController)
	p.PUT("/detail", controller.UpdateProfilDetailController)
	p.PUT("", controller.UpdateProfilController)
	p.DELETE("", controller.DeleteUserController)

	a := e.Group("/admin")
	a.Use(m.JWTMiddlewareConfig, m.IsAdmin)
	a.GET("", controller.DashboardAdminController)
	a.GET("/users", controller.GetUsersController)
	a.GET("/pets", controller.GetPetsAdminController)
	a.GET("/category", controller.GetPetCategoryController)
	a.POST("/category", controller.CreatePetCategoryController)
	// a.GET("/users/:username", controller.GetUserController)

	return e
}
