package routes

import (
	"BelajarAPI/controller/activity"

	"github.com/labstack/echo/v4"
)

func RouteActivity(e *echo.Echo, c *activity.ActivityController, jwtMiddleware echo.MiddlewareFunc) {
	e.POST("/activity", c.Add_Activity(), jwtMiddleware)
	e.PUT("/activity/:id", c.Edit_Activity(), jwtMiddleware)
	e.GET("/activities", c.Get_Activities(), jwtMiddleware)
}
