package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, c controller.UserController) {
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})

	e.POST("/login", c.Login())
	e.POST("/register", c.Register())
	// e.GET("/users", c.GetUsers(), jwtMiddleware)
	e.GET("/profile", c.GetUserByID(), jwtMiddleware)
	e.DELETE("/delete", c.DeleteUserByID(), jwtMiddleware)
	e.PUT("/update", c.UpdateUserByID(), jwtMiddleware)
}
