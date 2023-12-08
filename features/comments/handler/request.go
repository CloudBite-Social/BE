package handler

import "sosmed/features/comments"

type CreateCommentRequest struct {
	PostId uint   `json:"post_id"`
	Text   string `json:"text"`
}

func (req *CreateCommentRequest) ToEntity(userId uint) *comments.Comment {
	var ent = new(comments.Comment)

	if req.PostId != 0 {
		ent.PostId = req.PostId
	}

	if req.Text != "" {
		ent.Text = req.Text
	}

	if userId != 0 {
		ent.User.Id = userId
	}

	return ent
}
