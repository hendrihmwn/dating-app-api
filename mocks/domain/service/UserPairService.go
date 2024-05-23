// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	service "github.com/hendrihmwn/dating-app-api/app/domain/service"
	mock "github.com/stretchr/testify/mock"
)

// UserPairService is an autogenerated mock type for the UserPairService type
type UserPairService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, params
func (_m *UserPairService) Create(ctx context.Context, params service.CreateUserPairParams) (*service.CreateUserPairResponse, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *service.CreateUserPairResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, service.CreateUserPairParams) (*service.CreateUserPairResponse, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, service.CreateUserPairParams) *service.CreateUserPairResponse); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.CreateUserPairResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, service.CreateUserPairParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserPairService creates a new instance of UserPairService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserPairService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserPairService {
	mock := &UserPairService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
