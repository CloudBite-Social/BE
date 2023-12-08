package handler

import (
	"net/http"
	"sosmed/features/posts"
	"sosmed/features/users"
	"sosmed/helpers/tokens"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(RegisterUserRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		var data = request.ToEntity()

		if err := hdl.userService.Register(c.Request().Context(), *data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "Duplicate") {
				response["message"] = "email is already in use"
				return c.JSON(http.StatusConflict, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "register success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(LoginUserRequest)

		if err := c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		var input = request.ToEntity()

		result, token, err := hdl.userService.Login(c.Request().Context(), *input)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "wrong password") {
				response["message"] = err.Error()
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data = new(UserResponse)
		data.FromEntity(*result, token, nil)

		response["message"] = "login success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) GetById() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(UpdateUserRequest)

		token := c.Get("user")
		if token == nil {
			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		userId, err := tokens.ExtractToken(token.(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		if c.Bind(request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}

		file, _ := c.FormFile("image")
		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			request.Image = src
		}

		if err := hdl.userService.Update(c.Request().Context(), userId, *request.ToEntity()); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update user success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		userId, err := tokens.ExtractToken(c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		// TODO need post of user

		if err := hdl.userService.Delete(c.Request().Context(), userId); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "user not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete user success"
		return c.JSON(http.StatusOK, response)
	}
}
