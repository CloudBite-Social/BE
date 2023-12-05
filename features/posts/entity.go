package posts

import (
	"sosmed/features/comments"
	"sosmed/features/users"
	"time"
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
	Id   uint
	Path string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Handler interface{}

type Service interface{}

type Repository interface{}
