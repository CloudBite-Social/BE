package service

import (
	"context"
	"errors"
	"sosmed/features/users"
	"sosmed/helpers/encrypt"
)

func NewUserService(repo users.Repository, enc encrypt.BcryptHash) users.Service {
	return &userService{
		repo: repo,
		enc:  enc,
	}
}

type userService struct {
	repo users.Repository
	enc  encrypt.BcryptHash
}

func (srv *userService) Register(ctx context.Context, data users.User) error {
	if data.Name == "" {
		return errors.New("validate: name can't empty")
	}

	if data.Email == "" {
		return errors.New("validate: email can't empty")
	}

	if data.Password == "" {
		return errors.New("validate: password can't empty")
	}

	hash, err := srv.enc.Hash(data.Password)
	if err != nil {
		return err
	}

	data.Password = hash

	if err := srv.repo.Register(ctx, data); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Login(ctx context.Context, data users.User) (*users.User, *string, error) {
	panic("unimplemented")
}

func (srv *userService) GetById(ctx context.Context, id uint) (*users.User, error) {
	panic("unimplemented")
}

func (srv *userService) Update(ctx context.Context, id uint, data users.User) error {
	panic("unimplemented")
}

func (srv *userService) Delete(ctx context.Context, id uint) error {
	panic("unimplemented")
}
