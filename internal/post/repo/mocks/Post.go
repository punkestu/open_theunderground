// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/punkestu/open_theunderground/domain"
	mock "github.com/stretchr/testify/mock"
)

// Post is an autogenerated mock type for the Post type
type Post struct {
	mock.Mock
}

// Create provides a mock function with given fields: topic, authorId
func (_m *Post) Create(topic string, authorId string) (*domain.Post, error) {
	ret := _m.Called(topic, authorId)

	var r0 *domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.Post, error)); ok {
		return rf(topic, authorId)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.Post); ok {
		r0 = rf(topic, authorId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(topic, authorId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Post) GetAll() ([]*domain.Post, error) {
	ret := _m.Called()

	var r0 []*domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.Post, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAuthor provides a mock function with given fields: authorId
func (_m *Post) GetByAuthor(authorId string) ([]*domain.Post, error) {
	ret := _m.Called(authorId)

	var r0 []*domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*domain.Post, error)); ok {
		return rf(authorId)
	}
	if rf, ok := ret.Get(0).(func(string) []*domain.Post); ok {
		r0 = rf(authorId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(authorId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: postId
func (_m *Post) GetByID(postId string) (*domain.Post, error) {
	ret := _m.Called(postId)

	var r0 *domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Post, error)); ok {
		return rf(postId)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Post); ok {
		r0 = rf(postId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: topic
func (_m *Post) Update(topic string) (*domain.Post, error) {
	ret := _m.Called(topic)

	var r0 *domain.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Post, error)); ok {
		return rf(topic)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Post); ok {
		r0 = rf(topic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(topic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPost creates a new instance of Post. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPost(t interface {
	mock.TestingT
	Cleanup(func())
}) *Post {
	mock := &Post{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
