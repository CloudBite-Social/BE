package handler

import (
	"net/http"
	"sosmed/features/posts"
	"sosmed/helpers/tokens"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func NewPostHandler(service posts.Service) posts.Handler {
	return &postHandler{
		service: service,
	}
}

type postHandler struct {
	service posts.Service
}

func (hdl *postHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreatePostRequest)

		userID, err := tokens.ExtractToken(c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		request.Caption = c.FormValue("caption")

		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}
		files := form.File["image"]

		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				c.Logger().Error(err)

				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}
			defer src.Close()

			request.Files = append(request.Files, src)
		}

		data := request.ToEntity(userID)

		if err := hdl.service.Create(c.Request().Context(), *data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "invalid data") {
				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create post success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *postHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		postId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		result, err := hdl.service.GetById(c.Request().Context(), uint(postId))
		if err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "invalid data") {
				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		var data = new(PostResponse)
		data.FromEntity(*result)

		response["message"] = "get detail post success"
		response["data"] = data
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *postHandler) GetList() echo.HandlerFunc {
	panic("unimplemented")
}

func (hdl *postHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreatePostRequest)

		postId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		request.Caption = c.FormValue("caption")

		form, err := c.MultipartForm()
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}
		files := form.File["image"]

		for _, file := range files {
			src, err := file.Open()
			if err != nil {
				c.Logger().Error(err)

				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}
			defer src.Close()

			request.Files = append(request.Files, src)
		}

		data := request.ToEntity(0)

		if err := hdl.service.Update(c.Request().Context(), uint(postId), *data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "invalid data") {
				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "update post success"
		return c.JSON(http.StatusOK, response)
	}
}

func (hdl *postHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		postId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "bad request"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.service.Delete(c.Request().Context(), uint(postId)); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "invalid data") {
				response["message"] = "bad request"
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete post success"
		return c.JSON(http.StatusOK, response)
	}
}
