package handler

import (
	"io"
	"sosmed/features/posts"
)

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
