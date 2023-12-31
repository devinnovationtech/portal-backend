// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// TagRepository is an autogenerated mock type for the TagRepository type
type TagRepository struct {
	mock.Mock
}

// FetchTag provides a mock function with given fields: ctx, param
func (_m *TagRepository) FetchTag(ctx context.Context, param *domain.Request) ([]domain.Tag, int64, error) {
	ret := _m.Called(ctx, param)

	var r0 []domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) []domain.Tag); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Tag)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) int64); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.Request) error); ok {
		r2 = rf(ctx, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTagByName provides a mock function with given fields: ctx, name
func (_m *TagRepository) GetTagByName(ctx context.Context, name string) (domain.Tag, error) {
	ret := _m.Called(ctx, name)

	var r0 domain.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Tag); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(domain.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreTag provides a mock function with given fields: ctx, t, tx
func (_m *TagRepository) StoreTag(ctx context.Context, t *domain.Tag, tx *sql.Tx) error {
	ret := _m.Called(ctx, t, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Tag, *sql.Tx) error); ok {
		r0 = rf(ctx, t, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTagRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTagRepository creates a new instance of TagRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTagRepository(t mockConstructorTestingTNewTagRepository) *TagRepository {
	mock := &TagRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
