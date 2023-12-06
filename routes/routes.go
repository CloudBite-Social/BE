package routes

import (
	"sosmed/features/comments"
	"sosmed/features/posts"
	"sosmed/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Routes struct {
	Server *echo.Echo

	UserHandler    users.Handler
	PostHandler    posts.Handler
	CommentHandler comments.Handler
}

func (router Routes) InitRouter() {
	router.UserRouter()
	router.PostRouter()
	router.CommentRouter()
}

func (router *Routes) UserRouter() {
}

func (router *Routes) PostRouter() {
	router.Server.GET("/posts", router.PostHandler.GetList())
	router.Server.POST("/posts", router.PostHandler.Create(), echojwt.JWT([]byte("altamantul")))
	router.Server.GET("/posts/:id", router.PostHandler.GetById())
	router.Server.PUT("/posts/:id", router.PostHandler.Update(), echojwt.JWT([]byte("altamantul")))
	router.Server.DELETE("/posts/:id", router.PostHandler.Delete(), echojwt.JWT([]byte("altamantul")))
}

func (router *Routes) CommentRouter() {
}
