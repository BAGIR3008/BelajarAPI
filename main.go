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

	um := model.User(db)
	uc := controller.User(um)
	am := model.Activity(db)
	ac := controller.Activity(am)

	// config.Migrate(db, um.User)
	// config.Migrate(db, am.Activity)

	e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, uc, ac)
	e.Logger.Fatal(e.Start("127.0.0.1:8000"))
}
