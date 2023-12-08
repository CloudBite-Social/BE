package handler

import (
	"sosmed/features/posts"
	"sosmed/features/users"
	"time"
)

type AuthResponse struct {
	Id        uint      `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	Token string `json:"token,omitempty"`
}

func (res *AuthResponse) FromEntity(ent users.User, token *string) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.Email != "" {
		res.Email = ent.Email
	}

	if ent.Image != "" {
		res.Image = ent.Image
	} else {
		res.Image = "http://res.cloudinary.com/dxekaja1m/image/upload/v1701934913/users/aadvbkjit2nvzuw7ch4l.png"
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}

	if token != nil && *token != "" {
		res.Token = *token
	}
}

type UserResponse struct {
	Id        uint      `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	Posts []UserPostResponse `json:"posts"`
}

func (res *UserResponse) FromEntity(ent users.User, posts []posts.Post) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.Email != "" {
		res.Email = ent.Email
	}

	if ent.Image != "" {
		res.Image = ent.Image
	} else {
		res.Image = "http://res.cloudinary.com/dxekaja1m/image/upload/v1701934913/users/aadvbkjit2nvzuw7ch4l.png"
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}

	if len(posts) != 0 {
		for _, post := range posts {
			var tempPost = new(UserPostResponse)
			tempPost.FromEntity(post)

			res.Posts = append(res.Posts, *tempPost)
		}
	}
}

type UserPostResponse struct {
	Id      uint   `json:"post_id,omitempty"`
	Caption string `json:"caption,omitempty"`
	Image   string `json:"image,omitempty"`

	CommentCount *int      `json:"comment_count,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

func (res *UserPostResponse) FromEntity(ent posts.Post) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Caption != "" {
		res.Caption = ent.Caption
	}

	if len(ent.Attachment) > 0 {
		res.Image = ent.Attachment[0].URL
	}

	commentCount := len(ent.Comments)
	res.CommentCount = &commentCount

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}
}
