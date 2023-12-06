package service

import (
	"context"
	"errors"
	"sosmed/features/comments"
	"sosmed/features/comments/mocks"
	"sosmed/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentServiceCreate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewCommentService(repo)
	var ctx = context.Background()

	t.Run("invalid user id", func(t *testing.T) {
		caseInput := comments.Comment{
			Text:   "example comment 1",
			PostId: 1,
			User:   users.User{Id: 0},
		}

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("invalid post id", func(t *testing.T) {
		caseInput := comments.Comment{
			Text:   "example comment 1",
			PostId: 0,
			User:   users.User{Id: 1},
		}

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("invalid comment text", func(t *testing.T) {
		caseInput := comments.Comment{
			Text:   "",
			PostId: 1,
			User:   users.User{Id: 1},
		}

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("repository error", func(t *testing.T) {
		caseInput := comments.Comment{
			Text:   "example comment 1",
			PostId: 1,
			User:   users.User{Id: 1},
		}
		repo.On("Create", ctx, caseInput).Return(errors.New("some error from repository")).Once()

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseInput := comments.Comment{
			Text:   "example comment 1",
			PostId: 1,
			User:   users.User{Id: 1},
		}
		repo.On("Create", ctx, caseInput).Return(nil).Once()

		err := srv.Create(ctx, caseInput)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestCommentServiceDelete(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewCommentService(repo)
	var ctx = context.Background()

	t.Run("invalid comment id", func(t *testing.T) {
		err := srv.Delete(ctx, 0)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("repository error", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(errors.New("some error from repository")).Once()

		err := srv.Delete(ctx, 1)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(nil).Once()

		err := srv.Delete(ctx, 1)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}
