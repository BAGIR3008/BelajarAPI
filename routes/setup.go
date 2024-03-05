package routes

import (
	"BelajarAPI/config"
	"BelajarAPI/controller/activity"
	"BelajarAPI/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc *user.UserController, ac *activity.ActivityController) {
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})

	RouteUser(e, uc, jwtMiddleware)
	RouteActivity(e, ac, jwtMiddleware)
}
