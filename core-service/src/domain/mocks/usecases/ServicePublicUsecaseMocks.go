// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// ServicePublicUsecase is an autogenerated mock type for the ServicePublicUsecase type
type ServicePublicUsecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ServicePublicUsecase) Delete(_a0 context.Context, _a1 int64) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, params
func (_m *ServicePublicUsecase) Fetch(ctx context.Context, params *domain.Request) ([]domain.ServicePublic, error) {
	ret := _m.Called(ctx, params)

	var r0 []domain.ServicePublic
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) []domain.ServicePublic); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ServicePublic)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBySlug provides a mock function with given fields: ctx, slug
func (_m *ServicePublicUsecase) GetBySlug(ctx context.Context, slug string) (domain.ServicePublic, error) {
	ret := _m.Called(ctx, slug)

	var r0 domain.ServicePublic
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ServicePublic); ok {
		r0 = rf(ctx, slug)
	} else {
		r0 = ret.Get(0).(domain.ServicePublic)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, slug)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetaFetch provides a mock function with given fields: ctx, params
func (_m *ServicePublicUsecase) MetaFetch(ctx context.Context, params *domain.Request) (int64, string, int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) string); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 int64
	if rf, ok := ret.Get(2).(func(context.Context, *domain.Request) int64); ok {
		r2 = rf(ctx, params)
	} else {
		r2 = ret.Get(2).(int64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(context.Context, *domain.Request) error); ok {
		r3 = rf(ctx, params)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *ServicePublicUsecase) Store(_a0 context.Context, _a1 domain.StorePublicService) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.StorePublicService) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *ServicePublicUsecase) Update(_a0 context.Context, _a1 domain.UpdatePublicService, _a2 int64) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.UpdatePublicService, int64) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewServicePublicUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewServicePublicUsecase creates a new instance of ServicePublicUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServicePublicUsecase(t mockConstructorTestingTNewServicePublicUsecase) *ServicePublicUsecase {
	mock := &ServicePublicUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}