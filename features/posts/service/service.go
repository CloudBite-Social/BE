package service

import (
	"context"
	"errors"
	"sosmed/features/posts"
	"sosmed/helpers/filters"
)

func NewPostService(repo posts.Repository) posts.Service {
	return &postService{
		repo: repo,
	}
}

type postService struct {
	repo posts.Repository
}

func (srv *postService) Create(ctx context.Context, data posts.Post) error {
	if data.Caption == "" && len(data.Attachment) == 0 {
		return errors.New("validate: please fill image or caption")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *postService) GetById(ctx context.Context, postId uint) (*posts.Post, error) {
	if postId == 0 {
		return nil, errors.New("validate: invalid post id")
	}

	result, err := srv.repo.GetById(ctx, postId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *postService) GetList(ctx context.Context, filter filters.Filter, userId *uint) ([]posts.Post, int, error) {
	result, totalData, err := srv.repo.GetList(ctx, filter, userId)
	if err != nil {
		return nil, 0, err
	}

	return result, totalData, nil
}

func (srv *postService) Update(ctx context.Context, postId uint, data posts.Post) error {
	if postId == 0 {
		return errors.New("validate: invalid post id")
	}

	if data.Caption == "" && len(data.Attachment) == 0 {
		return errors.New("validate: please fill image or caption")
	}

	if err := srv.repo.Update(ctx, postId, data); err != nil {
		return err
	}

	return nil
}

func (srv *postService) Delete(ctx context.Context, postId uint) error {
	if postId == 0 {
		return errors.New("invalid data")
	}

	if err := srv.repo.Delete(ctx, postId); err != nil {
		return err
	}

	return nil
}
