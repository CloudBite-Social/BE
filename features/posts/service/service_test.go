package service

import (
	"context"
	"errors"
	"sosmed/features/posts"
	"sosmed/features/posts/mocks"
	"sosmed/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostServiceCreate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("invalid input data", func(t *testing.T) {
		caseInput := posts.Post{}

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("repository error", func(t *testing.T) {
		caseInput := posts.Post{
			Caption: "example caption 1",
			User:    users.User{Id: 1},
		}
		repo.On("Create", ctx, caseInput).Return(errors.New("some error from repository")).Once()

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseInput := posts.Post{
			Caption: "example caption 1",
			User:    users.User{Id: 1},
		}
		repo.On("Create", ctx, caseInput).Return(nil).Once()

		err := srv.Create(ctx, caseInput)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}
