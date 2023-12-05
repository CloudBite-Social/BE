package main

import (
	"sosmed/config"
	"sosmed/routes"
	"sosmed/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	var dbConfig = new(config.Database)
	if err := dbConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	dbConnection, err := database.MysqlInit(*dbConfig)
	if err != nil {
		panic(err)
	}

	if err := database.MysqlMigrate(dbConnection); err != nil {
		panic(err)
	}

	app := echo.New()

	route := routes.Routes{
		Server: app,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}
