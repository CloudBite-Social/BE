package comments

import (
	"sosmed/features/users"
	"time"
)

type Comment struct {
	Id   uint
	Text string

	PostId int

	User users.User

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface{}

type Service interface{}

type Repository interface{}
