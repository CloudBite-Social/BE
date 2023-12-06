package handler

import (
	"io"
	"reflect"
	"sosmed/features/posts"
	"time"
)

type PostResponse struct {
	Id      uint   `json:"id,omitempty"`
	Caption string `json:"caption,omitempty"`
	Image   string `json:"image,omitempty"`

	User struct {
		Id    uint   `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Image string `json:"image,omitempty"`

		CreatedAt time.Time `json:"created_at"`
	} `json:"user,omitempty"`

	CommentCount *int `json:"comment_count,omitempty"`
	Comments     []struct {
		Id        uint      `json:"id,omitempty"`
		Text      string    `json:"text,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		User      struct {
			Id    uint   `json:"id,omitempty"`
			Name  string `json:"name,omitempty"`
			Image string `json:"image,omitempty"`

			CreatedAt time.Time `json:"created_at"`
		} `json:"user,omitempty"`
	} `json:"comment"`

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
		if ent.User.Id != 0 {
			res.User.Id = ent.User.Id
		}

		if ent.User.Name != "" {
			res.User.Name = ent.User.Name
		}

		if ent.User.Image != "" {
			res.User.Image = ent.User.Image
		}

		if !ent.User.CreatedAt.IsZero() {
			res.User.CreatedAt = ent.User.CreatedAt
		}
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
				var tempComment = struct {
					Id        uint      `json:"id,omitempty"`
					Text      string    `json:"text,omitempty"`
					CreatedAt time.Time `json:"created_at"`
					User      struct {
						Id        uint      `json:"id,omitempty"`
						Name      string    `json:"name,omitempty"`
						Image     string    `json:"image,omitempty"`
						CreatedAt time.Time `json:"created_at"`
					} `json:"user,omitempty"`
				}{}

				if comment.Id != 0 {
					tempComment.Id = comment.Id
				}

				if comment.Text != "" {
					tempComment.Text = comment.Text
				}

				if !reflect.ValueOf(comment.User).IsZero() {
					if comment.User.Id != 0 {
						tempComment.User.Id = comment.User.Id
					}

					if comment.User.Name != "" {
						tempComment.User.Name = comment.User.Name
					}

					if comment.User.Image != "" {
						tempComment.User.Image = comment.User.Image
					}

					if !comment.User.CreatedAt.IsZero() {
						tempComment.User.CreatedAt = comment.User.CreatedAt
					}
				}

				if !comment.CreatedAt.IsZero() {
					tempComment.CreatedAt = comment.CreatedAt
				}

				res.Comments = append(res.Comments, tempComment)
			}
		}
	}

	if !ent.CreatedAt.IsZero() {
		res.CreatedAt = ent.CreatedAt
	}
}

type CreatePostRequest struct {
	Caption string
	Files   []io.Reader
}

func (req *CreatePostRequest) ToEntity(userId uint) *posts.Post {
	var ent = new(posts.Post)

	if userId != 0 {
		ent.User.Id = userId
	}

	if req.Caption != "" {
		ent.Caption = req.Caption
	}

	if len(req.Files) != 0 {
		for _, file := range req.Files {
			ent.Attachment = append(ent.Attachment, posts.File{Raw: file})
		}
	}
	return ent
}
