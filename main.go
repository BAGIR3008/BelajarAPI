package main

import (
	"BelajarAPI/config"
	ad "BelajarAPI/features/activity/data"
	ah "BelajarAPI/features/activity/handler"
	as "BelajarAPI/features/activity/services"
	"BelajarAPI/features/user/data"
	"BelajarAPI/features/user/handler"
	"BelajarAPI/features/user/services"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()            // inisiasi echo
	cfg := config.InitConfig() // baca seluruh system variable
	db := config.InitSQL(cfg)  // konek DB

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	activityData := ad.New(db)
	activityService := as.NewActivityService(activityData)
	activityHandler := ah.NewHandler(activityService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, userHandler, activityHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
