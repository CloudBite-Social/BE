package service

import (
	"context"
	"errors"
	"sosmed/features/comments"
	"sosmed/features/posts"
	"sosmed/features/posts/mocks"
	"sosmed/features/users"
	"sosmed/helpers/filters"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostServiceCreate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("invalid input data", func(t *testing.T) {
		caseInput := posts.Post{}

		err := srv.Create(ctx, caseInput)

		assert.ErrorContains(t, err, "validate")
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

func TestPostServiceGetById(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("invalid id", func(t *testing.T) {
		result, err := srv.GetById(ctx, 0)

		assert.ErrorContains(t, err, "validate")
		assert.Nil(t, result)
	})

	t.Run("repository error", func(t *testing.T) {
		repo.On("GetById", ctx, uint(1)).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.GetById(ctx, 1)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetById", ctx, uint(1)).Return(nil, errors.New("data not found")).Once()

		result, err := srv.GetById(ctx, 1)

		assert.ErrorContains(t, err, "data not found")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseResult := &posts.Post{
			Id:      1,
			Caption: "example caption 1",
			User: users.User{
				Id:        1,
				Name:      "kijang 1",
				Email:     "kijang1@mail.com",
				Password:  "$2a$10$dhhW17wM2yzPwD0qrURWHez5eUtyrYxFKSuqw/Udjpd22j1xzTP0W",
				Image:     "https://placehold.co/400x400/png",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Comments: []comments.Comment{
				{
					Id:     1,
					Text:   "example comment",
					PostId: 1,
					User: users.User{
						Id:        2,
						Name:      "kijang 2",
						Email:     "kijang2@mail.com",
						Password:  "$2a$10$xb6YMHB1G1.fQMtHahcDHuLf3b7E4pMdcCBPrYjpdrfO4ImUmQhhW",
						Image:     "https://placehold.co/400x400/png",
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			Attachment: []posts.File{
				{
					Id:        1,
					URL:       "https://placehold.co/600x400/png",
					Raw:       nil,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		repo.On("GetById", ctx, uint(1)).Return(caseResult, nil).Once()

		result, err := srv.GetById(ctx, 1)

		assert.NoError(t, err)
		assert.EqualValues(t, caseResult, result)

		repo.AssertExpectations(t)
	})
}

func TestPostServiceGetList(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("repository error", func(t *testing.T) {
		var caseFilter = filters.Filter{
			Search: filters.Search{Keyword: "caption"},
			Pagination: filters.Pagination{
				Limit: 1,
				Start: 0,
			},
		}
		repo.On("GetList", ctx, caseFilter, (*uint)(nil)).Return(nil, 0, errors.New("some error from repository")).Once()

		result, total, err := srv.GetList(ctx, caseFilter, nil)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Equal(t, 0, total)
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseTotal = 10
		var caseFilter = filters.Filter{
			Search: filters.Search{Keyword: "caption"},
			Pagination: filters.Pagination{
				Limit: 1,
				Start: 0,
			},
		}
		var caseResult = []posts.Post{
			{
				Id:      1,
				Caption: "example caption 1",
				User: users.User{
					Id:        1,
					Name:      "kijang 1",
					Email:     "kijang1@mail.com",
					Password:  "$2a$10$dhhW17wM2yzPwD0qrURWHez5eUtyrYxFKSuqw/Udjpd22j1xzTP0W",
					Image:     "https://placehold.co/400x400/png",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Comments: []comments.Comment{
					{
						Id:     1,
						Text:   "example comment",
						PostId: 1,
						User: users.User{
							Id:        2,
							Name:      "kijang 2",
							Email:     "kijang2@mail.com",
							Password:  "$2a$10$xb6YMHB1G1.fQMtHahcDHuLf3b7E4pMdcCBPrYjpdrfO4ImUmQhhW",
							Image:     "https://placehold.co/400x400/png",
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				Attachment: []posts.File{
					{
						Id:        1,
						URL:       "https://placehold.co/600x400/png",
						Raw:       nil,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					},
				},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		repo.On("GetList", ctx, caseFilter, (*uint)(nil)).Return(caseResult, caseTotal, nil).Once()

		result, total, err := srv.GetList(ctx, caseFilter, nil)

		assert.NoError(t, err)
		assert.Equal(t, caseTotal, total)
		assert.Equal(t, caseResult, result)

		repo.AssertExpectations(t)
	})
}

func TestPostServiceUpdate(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("invalid id", func(t *testing.T) {
		var caseData = posts.Post{}
		err := srv.Update(ctx, 0, caseData)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("invalid caption", func(t *testing.T) {
		var caseData = posts.Post{
			Caption: "",
		}
		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("invalid file attachment", func(t *testing.T) {
		var caseData = posts.Post{
			Caption:    "example post 1",
			Attachment: []posts.File{},
		}
		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("repository error", func(t *testing.T) {
		var caseData = posts.Post{
			Caption: "example post 1",
			Attachment: []posts.File{
				{
					Raw: nil,
				},
			},
		}
		repo.On("Update", ctx, uint(1), caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		var caseData = posts.Post{
			Caption: "example post 1",
			Attachment: []posts.File{
				{
					Raw: nil,
				},
			},
		}
		repo.On("Update", ctx, uint(1), caseData).Return(errors.New("data not found")).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.ErrorContains(t, err, "data not found")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = posts.Post{
			Caption: "example post 1",
			Attachment: []posts.File{
				{
					Raw: nil,
				},
			},
		}

		repo.On("Update", ctx, uint(1), caseData).Return(nil).Once()

		err := srv.Update(ctx, 1, caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestPostServiceDelete(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var srv = NewPostService(repo)
	var ctx = context.Background()

	t.Run("invalid id", func(t *testing.T) {
		err := srv.Delete(ctx, 0)

		assert.ErrorContains(t, err, "invalid data")
	})

	t.Run("repository error", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(errors.New("some error from repository")).Once()

		err := srv.Delete(ctx, 1)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(errors.New("data not found")).Once()

		err := srv.Delete(ctx, 1)

		assert.ErrorContains(t, err, "data not found")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("Delete", ctx, uint(1)).Return(nil).Once()

		err := srv.Delete(ctx, 1)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}
