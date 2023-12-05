package routes

import (
	"sosmed/features/comments"
	"sosmed/features/posts"
	"sosmed/features/users"

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
}

func (router *Routes) CommentRouter() {
}
