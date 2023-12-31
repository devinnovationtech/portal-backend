// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// PopUpBannerRepository is an autogenerated mock type for the PopUpBannerRepository type
type PopUpBannerRepository struct {
	mock.Mock
}

// DeactiveStatus provides a mock function with given fields: ctx
func (_m *PopUpBannerRepository) DeactiveStatus(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *PopUpBannerRepository) Delete(ctx context.Context, id int64) error {
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
func (_m *PopUpBannerRepository) Fetch(ctx context.Context, params *domain.Request) ([]domain.PopUpBanner, int64, error) {
	ret := _m.Called(ctx, params)

	var r0 []domain.PopUpBanner
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) []domain.PopUpBanner); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PopUpBanner)
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
func (_m *PopUpBannerRepository) GetByID(ctx context.Context, id int64) (domain.PopUpBanner, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.PopUpBanner
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.PopUpBanner); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.PopUpBanner)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LiveBanner provides a mock function with given fields: ctx
func (_m *PopUpBannerRepository) LiveBanner(ctx context.Context) (domain.PopUpBanner, error) {
	ret := _m.Called(ctx)

	var r0 domain.PopUpBanner
	if rf, ok := ret.Get(0).(func(context.Context) domain.PopUpBanner); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(domain.PopUpBanner)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, body
func (_m *PopUpBannerRepository) Store(ctx context.Context, body *domain.StorePopUpBannerRequest) error {
	ret := _m.Called(ctx, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.StorePopUpBannerRequest) error); ok {
		r0 = rf(ctx, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, id, body
func (_m *PopUpBannerRepository) Update(ctx context.Context, id int64, body *domain.StorePopUpBannerRequest) error {
	ret := _m.Called(ctx, id, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.StorePopUpBannerRequest) error); ok {
		r0 = rf(ctx, id, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatus provides a mock function with given fields: ctx, id, body
func (_m *PopUpBannerRepository) UpdateStatus(ctx context.Context, id int64, body *domain.UpdateStatusPopUpBannerRequest) error {
	ret := _m.Called(ctx, id, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *domain.UpdateStatusPopUpBannerRequest) error); ok {
		r0 = rf(ctx, id, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPopUpBannerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPopUpBannerRepository creates a new instance of PopUpBannerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPopUpBannerRepository(t mockConstructorTestingTNewPopUpBannerRepository) *PopUpBannerRepository {
	mock := &PopUpBannerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
