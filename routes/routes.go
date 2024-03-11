package routes

import (
	"BelajarAPI/config"
	activity "BelajarAPI/features/activity"
	user "BelajarAPI/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, uc user.UserController, ac activity.ActivityController) {
	config := echojwt.WithConfig(echojwt.Config{SigningKey: []byte(config.JWTSECRET)})

	userRoute(c, uc, config)
	activityRoute(c, ac, config)
}

func userRoute(c *echo.Echo, uc user.UserController, config echo.MiddlewareFunc) {
	c.POST("/register", uc.Register())
	c.POST("/login", uc.Login())
	c.GET("/profile", uc.Profile(), config)
}

func activityRoute(c *echo.Echo, ac activity.ActivityController, config echo.MiddlewareFunc) {
	c.POST("/activity", ac.Add(), config)
	c.PUT("/activity/:id", ac.Update(), config)
	c.GET("/activity", ac.GetAll(), config)
	c.DELETE("/activity/:id", ac.Delete(), config)
}
