package users

import (
	"context"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string

	Image    string
	RawImage io.Reader

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Register(ctx context.Context, data User) error
	Login(ctx context.Context, data User) (*User, *string, error)
	GetById(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, id uint, data User) error
	Delete(ctx context.Context, id uint) error
}

type Repository interface {
	Register(ctx context.Context, data User) error
	Login(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, id uint, data User) error
	Delete(ctx context.Context, id uint) error
}
