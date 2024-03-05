package main

import (
	"BelajarAPI/config"
	"BelajarAPI/controller"
	"BelajarAPI/model"
	"BelajarAPI/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userModel := model.User(db)
	userController := controller.User(userModel)
	activityModel := model.Activity(db)
	activityController := controller.Activity(activityModel)
	config.Migrate(db, userModel.User)
	config.Migrate(db, activityModel.Activity)

	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userController, activityController)
	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
