package handler

import (
	"net/http"
	"sosmed/features/comments"
	"sosmed/helpers/tokens"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echo "github.com/labstack/echo/v4"
)

func NewCommentHandler(service comments.Service) comments.Handler {
	return &commentHandler{
		service: service,
	}
}

type commentHandler struct {
	service comments.Service
}

func (hdl *commentHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)
		var request = new(CreateCommentRequest)

		userID, err := tokens.ExtractToken(c.Get("user").(*jwt.Token))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "unauthorized"
			return c.JSON(http.StatusUnauthorized, response)
		}

		if err := c.Bind(&request); err != nil {
			c.Logger().Error(err)

			response["message"] = "please fill input correctly"
			return c.JSON(http.StatusBadRequest, response)
		}
		data := request.ToEntity(userID)

		if err := hdl.service.Create(c.Request().Context(), *data); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = err.Error()
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "create comment success"
		return c.JSON(http.StatusCreated, response)
	}
}

func (hdl *commentHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response = make(map[string]any)

		commentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Logger().Error(err)

			response["message"] = "invalid comment id"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := hdl.service.Delete(c.Request().Context(), uint(commentId)); err != nil {
			c.Logger().Error(err)

			if strings.Contains(err.Error(), "validate: ") {
				response["message"] = strings.ReplaceAll(err.Error(), "validate: ", "")
				return c.JSON(http.StatusBadRequest, response)
			}

			if strings.Contains(err.Error(), "not found") {
				response["message"] = "comment not found"
				return c.JSON(http.StatusNotFound, response)
			}

			response["message"] = "internal server error"
			return c.JSON(http.StatusInternalServerError, response)
		}

		response["message"] = "delete comment success"
		return c.JSON(http.StatusOK, response)
	}
}
