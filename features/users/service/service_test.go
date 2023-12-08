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

func TestUserServiceLogin(t *testing.T) {}

func TestUserServiceGetById(t *testing.T) {}

func TestUserServiceUpdate(t *testing.T) {}

func TestUserServiceDelete(t *testing.T) {}
