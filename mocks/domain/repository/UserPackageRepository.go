// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/hendrihmwn/dating-app-api/app/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserPackageRepository is an autogenerated mock type for the UserPackageRepository type
type UserPackageRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, userPackage
func (_m *UserPackageRepository) Create(ctx context.Context, userPackage *entity.UserPackage) (*entity.UserPackage, error) {
	ret := _m.Called(ctx, userPackage)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entity.UserPackage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserPackage) (*entity.UserPackage, error)); ok {
		return rf(ctx, userPackage)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.UserPackage) *entity.UserPackage); ok {
		r0 = rf(ctx, userPackage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserPackage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.UserPackage) error); ok {
		r1 = rf(ctx, userPackage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, userId
func (_m *UserPackageRepository) Get(ctx context.Context, userId int64) (*entity.UserPackage, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.UserPackage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.UserPackage, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.UserPackage); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserPackage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, userId
func (_m *UserPackageRepository) List(ctx context.Context, userId int64) ([]entity.UserPackage, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []entity.UserPackage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) ([]entity.UserPackage, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) []entity.UserPackage); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UserPackage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserPackageRepository creates a new instance of UserPackageRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserPackageRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserPackageRepository {
	mock := &UserPackageRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
