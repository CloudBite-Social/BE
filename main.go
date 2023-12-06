package main

import (
	"sosmed/config"
	"sosmed/routes"
	"sosmed/utils/database"

	ph "sosmed/features/posts/handler"
	pr "sosmed/features/posts/repository"
	ps "sosmed/features/posts/service"

	"github.com/cloudinary/cloudinary-go/v2"
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

	var cldConfig = new(config.Cloudinary)
	if err := cldConfig.LoadFromEnv(); err != nil {
		panic(err)
	}

	cld, err := cloudinary.NewFromParams(cldConfig.CloudName, cldConfig.ApiKey, cldConfig.ApiSecret)
	if err != nil {
		panic(err)
	}

	postRepository := pr.NewPostRepository(dbConnection, cld)
	postService := ps.NewPostService(postRepository)
	postHandler := ph.NewPostHandler(postService)

	app := echo.New()

	route := routes.Routes{
		Server:      app,
		PostHandler: postHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}
