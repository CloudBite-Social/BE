package handler

import (
	"io"
	"reflect"
	"sosmed/features/posts"
)

type PostResponse struct {
	Id      uint   `json:"id,omitempty"`
	Caption string `json:"caption,omitempty"`
	Image   string `json:"image,omitempty"`

	User struct {
		Id    uint   `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Image string `json:"image,omitempty"`
	} `json:"user,omitempty"`

	Comments []struct {
		Id   uint   `json:"id,omitempty"`
		Text string `json:"text,omitempty"`
		User struct {
			Id    uint   `json:"id,omitempty"`
			Name  string `json:"name,omitempty"`
			Image string `json:"image,omitempty"`
		} `json:"user,omitempty"`
	} `json:"comment,omitempty"`
}

func (res *PostResponse) FromEntity(ent posts.Post) {
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
	}

	if len(ent.Comments) != 0 {
		for _, comment := range ent.Comments {
			if !reflect.ValueOf(comment).IsZero() {
				var tempComment = struct {
					Id   uint   `json:"id,omitempty"`
					Text string `json:"text,omitempty"`
					User struct {
						Id    uint   `json:"id,omitempty"`
						Name  string `json:"name,omitempty"`
						Image string `json:"image,omitempty"`
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
				}

				res.Comments = append(res.Comments, tempComment)
			}
		}
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
