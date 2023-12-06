package service

import (
	"context"
	"errors"
	"reflect"
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
	if reflect.ValueOf(data).IsZero() {
		return errors.New("invalid data")
	}

	if err := srv.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *postService) GetById(ctx context.Context, postId uint) (*posts.Post, error) {
	panic("unimplemented")
}

func (srv *postService) GetList(ctx context.Context, filter filters.Filter) ([]posts.Post, error) {
	panic("unimplemented")
}

func (srv *postService) Update(ctx context.Context, postId uint, data posts.Post) error {
	panic("unimplemented")
}

func (srv *postService) Delete(ctx context.Context, postId uint) error {
	panic("unimplemented")
}
