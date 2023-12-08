package service

import (
	"context"
	"errors"
	"sosmed/features/users"
	"sosmed/features/users/mocks"
	encMock "sosmed/helpers/encrypt/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceRegister(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewBcryptHash(t)
	var srv = NewUserService(repo, enc)
	var ctx = context.Background()

	t.Run("invalid name", func(t *testing.T) {
		var caseData = users.User{
			Name:     "",
			Email:    "kijang1@mail.com",
			Password: "test",
		}

		err := srv.Register(ctx, caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Name:     "kijang 1",
			Email:    "",
			Password: "test",
		}

		err := srv.Register(ctx, caseData)

		assert.ErrorContains(t, err, "email")
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Password: "",
		}

		err := srv.Register(ctx, caseData)

		assert.ErrorContains(t, err, "password")
	})

	t.Run("error from encrypt", func(t *testing.T) {
		var caseData = users.User{
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("", errors.New("some error from encrypt")).Once()

		err := srv.Register(ctx, caseData)

		assert.ErrorContains(t, err, "some error from encrypt")

		enc.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", ctx, caseData).Return(errors.New("some error from repository")).Once()

		caseData.Password = "test"
		err := srv.Register(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", ctx, caseData).Return(nil).Once()

		caseData.Password = "test"
		err := srv.Register(ctx, caseData)

		assert.NoError(t, err)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}

func TestUserServiceLogin(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewBcryptHash(t)
	var srv = NewUserService(repo, enc)
	var ctx = context.Background()

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Email:    "",
			Password: "kijang1",
		}

		result, token, err := srv.Login(ctx, caseData)

		assert.ErrorContains(t, err, "email")
		assert.Nil(t, result)
		assert.Nil(t, token)
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Email:    "kijang1@mail.com",
			Password: "",
		}

		result, token, err := srv.Login(ctx, caseData)

		assert.ErrorContains(t, err, "password")
		assert.Nil(t, result)
		assert.Nil(t, token)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Email:    "kijang1@mail.com",
			Password: "kijang1",
		}

		repo.On("Login", ctx, caseData.Email).Return(nil, errors.New("some error from repository")).Once()

		result, token, err := srv.Login(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)
		assert.Nil(t, token)

		repo.AssertExpectations(t)
	})

	t.Run("error from encrypt", func(t *testing.T) {
		var caseData = users.User{
			Email:    "kijang1@mail.com",
			Password: "kijang1",
		}

		var caseResult = users.User{
			Id:       1,
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Image:    "https://placehold.co/400x400/png",
			Password: "$2a$10$lT5fLMaj8497a1DBntlX5eMQi7/rkaV66JipGX80VrBBfLTxYv6GS",
		}

		repo.On("Login", ctx, caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(errors.New("some error from encrypt")).Once()

		result, token, err := srv.Login(ctx, caseData)

		assert.ErrorContains(t, err, "wrong password")
		assert.Nil(t, result)
		assert.Nil(t, token)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("error invalid id from generate jwt", func(t *testing.T) {
		var caseData = users.User{
			Email:    "kijang1@mail.com",
			Password: "kijang1",
		}

		var caseResult = users.User{
			Id:       0,
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Image:    "https://placehold.co/400x400/png",
			Password: "$2a$10$lT5fLMaj8497a1DBntlX5eMQi7/rkaV66JipGX80VrBBfLTxYv6GS",
		}

		repo.On("Login", ctx, caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(nil).Once()

		result, token, err := srv.Login(ctx, caseData)

		assert.ErrorContains(t, err, "invalid user id")
		assert.Nil(t, result)
		assert.Nil(t, token)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Email:    "kijang1@mail.com",
			Password: "kijang1",
		}

		var caseResult = users.User{
			Id:       1,
			Name:     "kijang 1",
			Email:    "kijang1@mail.com",
			Image:    "https://placehold.co/400x400/png",
			Password: "$2a$10$lT5fLMaj8497a1DBntlX5eMQi7/rkaV66JipGX80VrBBfLTxYv6GS",
		}

		repo.On("Login", ctx, caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(nil).Once()

		result, token, err := srv.Login(ctx, caseData)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, token)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}

func TestUserServiceGetById(t *testing.T) {}

func TestUserServiceUpdate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewBcryptHash(t)
	var srv = NewUserService(repo, enc)
	var ctx = context.Background()

	t.Run("invalid entity", func(t *testing.T) {
		caseData := users.User{}

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "please fill input correctly")
	})

	t.Run("invalid id", func(t *testing.T) {
		caseData := users.User{
			Name: "kijang 1",
		}

		err := srv.Update(ctx, 0, caseData)

		assert.ErrorContains(t, err, "invalid user id")
	})

	t.Run("error from encrypt", func(t *testing.T) {
		var caseData = users.User{
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("", errors.New("some error from encrypt")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "some error from encrypt")

		enc.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := users.User{
			Name: "kijang 1",
		}

		repo.On("Update", ctx, uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := users.User{
			Name:     "kijang 1",
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Update", ctx, uint(1), caseData).Return(nil).Once()

		caseData.Password = "test"
		err := srv.Update(ctx, 1, caseData)
		assert.Nil(t, err)

		repo.AssertExpectations(t)
	})
}

func TestUserServiceDelete(t *testing.T) {}
