// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
	mock "github.com/stretchr/testify/mock"
)

// GeneralInformationUsecase is an autogenerated mock type for the GeneralInformationUsecase type
type GeneralInformationUsecase struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *GeneralInformationUsecase) GetByID(ctx context.Context, id int64) (domain.GeneralInformation, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.GeneralInformation
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.GeneralInformation); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.GeneralInformation)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGeneralInformationUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewGeneralInformationUsecase creates a new instance of GeneralInformationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGeneralInformationUsecase(t mockConstructorTestingTNewGeneralInformationUsecase) *GeneralInformationUsecase {
	mock := &GeneralInformationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}