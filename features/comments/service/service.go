package service

import (
	"context"
	"errors"
	"sosmed/features/comments"
)

func NewCommentService(repo comments.Repository) comments.Service {
	return &commentService{
		repo: repo,
	}
}

type commentService struct {
	repo comments.Repository
}

func (srv *commentService) Create(ctx context.Context, data comments.Comment) error {
	if data.User.Id == 0 {
		return errors.New("validate: invalid user id")
	}

	if data.PostId == 0 {
		return errors.New("validate: invalid post id")
	}

	if data.Text == "" {
		return errors.New("validate: invalid text comment")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *commentService) Delete(ctx context.Context, commentId uint) error {
	if commentId == 0 {
		return errors.New("invalid data")
	}

	if err := srv.repo.Delete(ctx, commentId); err != nil {
		return err
	}

	return nil
}
