// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// NewsRepository is an autogenerated mock type for the NewsRepository type
type NewsRepository struct {
	mock.Mock
}

// AddView provides a mock function with given fields: ctx, id
func (_m *NewsRepository) AddView(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, params
func (_m *NewsRepository) Fetch(ctx context.Context, params *domain.Request) ([]domain.News, int64, error) {
	ret := _m.Called(ctx, params)

	var r0 []domain.News
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) []domain.News); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.News)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) int64); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *domain.Request) error); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *NewsRepository) GetByID(ctx context.Context, id int64) (domain.News, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.News
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.News); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.News)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
