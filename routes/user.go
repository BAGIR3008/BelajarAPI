package routes

import (
	"BelajarAPI/controller/user"

	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, c *user.UserController, jwtMiddleware echo.MiddlewareFunc) {
	e.POST("/login", c.Login())
	e.POST("/register", c.Register())
	e.GET("/profile", c.Profile(), jwtMiddleware)
}
