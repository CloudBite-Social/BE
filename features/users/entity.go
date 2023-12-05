package users

import (
	"time"
)

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
	Image    string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface{}

type Service interface{}

type Repository interface{}
