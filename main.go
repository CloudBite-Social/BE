package main

import (
	"sosmed/config"
	"sosmed/helpers/encrypt"
	"sosmed/routes"
	"sosmed/utils/database"

	uh "sosmed/features/users/handler"
	ur "sosmed/features/users/repository"
	us "sosmed/features/users/service"

	ph "sosmed/features/posts/handler"
	pr "sosmed/features/posts/repository"
	ps "sosmed/features/posts/service"

	ch "sosmed/features/comments/handler"
	cr "sosmed/features/comments/repository"
	cs "sosmed/features/comments/service"

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

	enc := encrypt.NewBcrypt(10)

	postRepository := pr.NewPostRepository(dbConnection, cld)
	postService := ps.NewPostService(postRepository)
	postHandler := ph.NewPostHandler(postService)

	userRepository := ur.NewUserRepository(dbConnection, cld)
	userService := us.NewUserService(userRepository, enc)
	userHandler := uh.NewUserHandler(userService, postService)

	commentRepository := cr.NewCommentRepository(dbConnection)
	commentService := cs.NewCommentService(commentRepository)
	commentHandler := ch.NewCommentHandler(commentService)

	app := echo.New()

	route := routes.Routes{
		Server:         app,
		PostHandler:    postHandler,
		CommentHandler: commentHandler,
		UserHandler:    userHandler,
	}

	route.InitRouter()

	app.Logger.Fatal(app.Start(":8000"))
}
