package handler

import (
	"fmt"
	"net/http"
	"sosmed/features/comments"
	"sosmed/features/posts"
	"sosmed/features/users"
	"sosmed/helpers/filters"
	"sosmed/helpers/tokens"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func NewUserHandler(userService users.Service, postService posts.Service, commetService comments.Service) users.Handler {
	return &userHandler{
		userService:   userService,
		postService:   postService,
		commetService: commetService,
	}
}

type userHandler struct {
	userService   users.Service
	postService   posts.Service
	commetService comments.Service
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

		var data = new(AuthResponse)
		data.FromEntity(*result, token)

		response["message"] = "login success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *userHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var baseUrl = c.Scheme() + "://" + c.Request().Host

		var pagination = new(filters.Pagination)
		c.Bind(pagination)
		if pagination.Limit == 0 {
			pagination.Limit = 5
		}

		var search = new(filters.Search)
		c.Bind(search)

		userId, err := tokens.ExtractToken(c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		resultUser, err := hdl.userService.GetById(c.Request().Context(), userId)
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		resultPost, totalData, err := hdl.postService.GetList(c.Request().Context(), filters.Filter{Pagination: *pagination, Search: *search}, &userId)
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var paginationResponse = make(map[string]any)
		if pagination.Start >= (pagination.Limit) {
			prev := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start-pagination.Limit, pagination.Limit)
			if search.Keyword != "" {
				prev += "&keyword=" + search.Keyword
			}
			paginationResponse["prev"] = prev
		} else {
			paginationResponse["prev"] = nil
		}

		if totalData > pagination.Start+pagination.Limit {
			next := fmt.Sprintf("%s%s?start=%d&limit=%d", baseUrl, c.Path(), pagination.Start+pagination.Limit, pagination.Limit)
			if search.Keyword != "" {
				next += "&keyword=" + search.Keyword
			}
			paginationResponse["next"] = next
		} else {
			paginationResponse["next"] = nil
		}

		var data = new(UserResponse)
		data.FromEntity(*resultUser, resultPost)

		response["message"] = "get detail of user success"
		response["pagination"] = paginationResponse
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
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

		if err := hdl.postService.DeleteByUserId(c.Request().Context(), userId); err != nil {
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

		if err := hdl.commetService.DeleteByUserId(c.Request().Context(), userId); err != nil {
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
