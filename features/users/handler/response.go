package handler

import (
	"time"
)

type UserResponse struct {
	Id        uint      `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	Posts []UserPostResponse `json:"posts"`

	Token string `json:"token,omitempty"`
}

type UserPostResponse struct {
	Id      uint   `json:"post_id,omitempty"`
	Caption string `json:"caption,omitempty"`
	Image   string `json:"image,omitempty"`

	CommentCount *int      `json:"comment_count,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
