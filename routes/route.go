package route

import (
	m "mini-project/middlewares"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	m.LoggerMiddleware(e)
	e.Pre(middleware.RemoveTrailingSlash())

	// e.POST("/register")
	// e.POST("/login")

	return e
}
