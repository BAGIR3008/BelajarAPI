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
	e.GET("/profile", c.Profile(), jwtMiddleware)
	e.POST("/activity", c.Add_Activity(), jwtMiddleware)
	e.PUT("/activity/:id", c.Edit_Activity(), jwtMiddleware)
	e.GET("/activities", c.Get_Activities(), jwtMiddleware)
}
