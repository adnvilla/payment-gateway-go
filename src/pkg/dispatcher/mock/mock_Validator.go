// Code generated by mockery v2.42.3. DO NOT EDIT.

package mock

import (
	context "context"

	dispatcher "github.com/adnvilla/payment-gateway-go/src/pkg/dispatcher"
	mock "github.com/stretchr/testify/mock"
)

// MockValidator is an autogenerated mock type for the Validator type
type MockValidator[TRequest dispatcher.Request] struct {
	mock.Mock
}

type MockValidator_Expecter[TRequest dispatcher.Request] struct {
	mock *mock.Mock
}

func (_m *MockValidator[TRequest]) EXPECT() *MockValidator_Expecter[TRequest] {
	return &MockValidator_Expecter[TRequest]{mock: &_m.Mock}
}

// Validate provides a mock function with given fields: ctx, request
func (_m *MockValidator[TRequest]) Validate(ctx context.Context, request TRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, TRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockValidator_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type MockValidator_Validate_Call[TRequest dispatcher.Request] struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - ctx context.Context
//   - request TRequest
func (_e *MockValidator_Expecter[TRequest]) Validate(ctx interface{}, request interface{}) *MockValidator_Validate_Call[TRequest] {
	return &MockValidator_Validate_Call[TRequest]{Call: _e.mock.On("Validate", ctx, request)}
}

func (_c *MockValidator_Validate_Call[TRequest]) Run(run func(ctx context.Context, request TRequest)) *MockValidator_Validate_Call[TRequest] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(TRequest))
	})
	return _c
}

func (_c *MockValidator_Validate_Call[TRequest]) Return(_a0 error) *MockValidator_Validate_Call[TRequest] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockValidator_Validate_Call[TRequest]) RunAndReturn(run func(context.Context, TRequest) error) *MockValidator_Validate_Call[TRequest] {
	_c.Call.Return(run)
	return _c
}

// NewMockValidator creates a new instance of MockValidator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockValidator[TRequest dispatcher.Request](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockValidator[TRequest] {
	mock := &MockValidator[TRequest]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
