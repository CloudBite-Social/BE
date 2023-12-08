package handler

import (
	"sosmed/features/posts"
	"sosmed/features/users"

	echo "github.com/labstack/echo/v4"
)

func NewUserHandler(userService users.Service, postService posts.Service) users.Handler {
	return &userHandler{
		userService: userService,
		postService: postService,
	}
}

type userHandler struct {
	userService users.Service
	postService posts.Service
}

func (hdl *userHandler) Register() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Login() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) GetById() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Update() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Delete() echo.HandlerFunc {
	panic("unimplemented")
}
