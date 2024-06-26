// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// PackageHandler is an autogenerated mock type for the PackageHandler type
type PackageHandler struct {
	mock.Mock
}

// Get provides a mock function with given fields: c
func (_m *PackageHandler) Get(c *gin.Context) {
	_m.Called(c)
}

// List provides a mock function with given fields: c
func (_m *PackageHandler) List(c *gin.Context) {
	_m.Called(c)
}

// NewPackageHandler creates a new instance of PackageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPackageHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *PackageHandler {
	mock := &PackageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
