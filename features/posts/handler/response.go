package handler

import (
	"reflect"
	"sosmed/features/comments"
	"sosmed/features/posts"
	"sosmed/features/users"
	"time"
)

type PostResponse struct {
	Id      uint   `json:"post_id,omitempty"`
	Caption string `json:"caption,omitempty"`
	Image   string `json:"image,omitempty"`

	User PostUserResponse `json:"user,omitempty"`

	CommentCount *int                  `json:"comment_count,omitempty"`
	Comments     []PostCommentResponse `json:"comment"`

	CreatedAt time.Time `json:"created_at"`
}

func (res *PostResponse) FromEntity(ent posts.Post, onlyCommentCount bool) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Caption != "" {
		res.Caption = ent.Caption
	}

	if len(ent.Attachment) > 0 {
		res.Image = ent.Attachment[0].URL
	}

	if !reflect.ValueOf(ent.User).IsZero() {
		res.User.FromEntity(ent.User)
	}

	if onlyCommentCount {
		commentCount := 0
		if len(ent.Comments) != 0 {
			commentCount = len(ent.Comments)
		}
		res.CommentCount = &commentCount
	} else if len(ent.Comments) != 0 {
		for _, comment := range ent.Comments {
			if !reflect.ValueOf(comment).IsZero() {
				var tempComment = new(PostCommentResponse)
				tempComment.FromEntity(comment)
				res.Comments = append(res.Comments, *tempComment)
			}
		}
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}
}

type PostUserResponse struct {
	Id    uint   `json:"user_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

func (res *PostUserResponse) FromEntity(ent users.User) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Name != "" {
		res.Name = ent.Name
	}

	if ent.Image != "" {
		res.Image = ent.Image
	} else {
		res.Image = "http://res.cloudinary.com/dxekaja1m/image/upload/v1701934913/users/aadvbkjit2nvzuw7ch4l.png"
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}
}

type PostCommentResponse struct {
	Id        uint             `json:"comment_id,omitempty"`
	Text      string           `json:"text,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	User      PostUserResponse `json:"user,omitempty"`
}

func (res *PostCommentResponse) FromEntity(ent comments.Comment) {
	if ent.Id != 0 {
		res.Id = ent.Id
	}

	if ent.Text != "" {
		res.Text = ent.Text
	}

	if !reflect.ValueOf(ent.User).IsZero() {
		res.User.FromEntity(ent.User)
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}
}
