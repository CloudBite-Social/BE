package service

import (
	"context"
	"errors"
	"sosmed/features/users"
	"sosmed/helpers/encrypt"
	"sosmed/helpers/tokens"
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
	if data.Email == "" {
		return nil, nil, errors.New("validate: email can't empty")
	}

	if data.Password == "" {
		return nil, nil, errors.New("validate: password can't empty")
	}

	result, err := srv.repo.Login(ctx, data.Email)
	if err != nil {
		return nil, nil, err
	}

	if err := srv.enc.Compare(result.Password, data.Password); err != nil {
		return nil, nil, errors.New("validate: wrong password")
	}

	token, err := tokens.GenerateJWT(result.Id)
	if err != nil {
		return nil, nil, err
	}

	return result, &token, nil
}

func (srv *userService) GetById(ctx context.Context, id uint) (*users.User, error) {
	if id == 0 {
		return nil, errors.New("validate: invalid user id")
	}

	result, err := srv.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (srv *userService) Update(ctx context.Context, id uint, data users.User) error {
	if id == 0 {
		return errors.New("validate: invalid user id")
	}

	if data.Password != "" {
		hash, err := srv.enc.Hash(data.Password)
		if err != nil {
			return err
		}

		data.Password = hash
	}

	if err := srv.repo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}

func (srv *userService) Delete(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("validate: invalid user id")
	}

	if err := srv.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
