package posts

import (
	"context"
	"io"
	"sosmed/features/comments"
	"sosmed/features/users"
	"sosmed/helpers/filters"
	"time"

	"github.com/labstack/echo/v4"
)

type Post struct {
	Id      uint
	Caption string

	User     users.User
	Comments []comments.Comment

	Attachment []File

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type File struct {
	Id  uint
	URL string

	Raw io.Reader

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	GetList() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	GetList(ctx context.Context, filter filters.Filter, userId *uint) ([]Post, int, error)
	GetById(ctx context.Context, postId uint) (*Post, error)
	Create(ctx context.Context, data Post) error
	Update(ctx context.Context, postId uint, data Post) error
	Delete(ctx context.Context, postId uint) error
}

type Repository interface {
	GetList(ctx context.Context, filter filters.Filter, userId *uint) ([]Post, int, error)
	GetById(ctx context.Context, postId uint) (*Post, error)
	Create(ctx context.Context, data Post) error
	Update(ctx context.Context, postId uint, data Post) error
	Delete(ctx context.Context, postId uint) error
}
