// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	filters "sosmed/helpers/filters"

	mock "github.com/stretchr/testify/mock"

	posts "sosmed/features/posts"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Service) Create(ctx context.Context, data posts.Post) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, posts.Post) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, postId
func (_m *Service) Delete(ctx context.Context, postId uint) error {
	ret := _m.Called(ctx, postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: ctx, postId
func (_m *Service) GetById(ctx context.Context, postId uint) (*posts.Post, error) {
	ret := _m.Called(ctx, postId)

	var r0 *posts.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) (*posts.Post, error)); ok {
		return rf(ctx, postId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint) *posts.Post); ok {
		r0 = rf(ctx, postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*posts.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(ctx, postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields: ctx, filter
func (_m *Service) GetList(ctx context.Context, filter filters.Filter) ([]posts.Post, int, error) {
	ret := _m.Called(ctx, filter)

	var r0 []posts.Post
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) ([]posts.Post, int, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, filters.Filter) []posts.Post); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]posts.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, filters.Filter) int); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, filters.Filter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, postId, data
func (_m *Service) Update(ctx context.Context, postId uint, data posts.Post) error {
	ret := _m.Called(ctx, postId, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint, posts.Post) error); ok {
		r0 = rf(ctx, postId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
