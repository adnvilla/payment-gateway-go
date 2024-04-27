// Code generated by mockery v2.42.3. DO NOT EDIT.

package mock

import (
	context "context"

	usecase "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	mock "github.com/stretchr/testify/mock"
)

// MockCreateOrderUseCase is an autogenerated mock type for the CreateOrderUseCase type
type MockCreateOrderUseCase struct {
	mock.Mock
}

type MockCreateOrderUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreateOrderUseCase) EXPECT() *MockCreateOrderUseCase_Expecter {
	return &MockCreateOrderUseCase_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, input
func (_m *MockCreateOrderUseCase) Handle(ctx context.Context, input usecase.CreateOrderInput) (usecase.CreateOrderOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 usecase.CreateOrderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.CreateOrderInput) (usecase.CreateOrderOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.CreateOrderInput) usecase.CreateOrderOutput); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(usecase.CreateOrderOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.CreateOrderInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCreateOrderUseCase_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockCreateOrderUseCase_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - input usecase.CreateOrderInput
func (_e *MockCreateOrderUseCase_Expecter) Handle(ctx interface{}, input interface{}) *MockCreateOrderUseCase_Handle_Call {
	return &MockCreateOrderUseCase_Handle_Call{Call: _e.mock.On("Handle", ctx, input)}
}

func (_c *MockCreateOrderUseCase_Handle_Call) Run(run func(ctx context.Context, input usecase.CreateOrderInput)) *MockCreateOrderUseCase_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(usecase.CreateOrderInput))
	})
	return _c
}

func (_c *MockCreateOrderUseCase_Handle_Call) Return(_a0 usecase.CreateOrderOutput, _a1 error) *MockCreateOrderUseCase_Handle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCreateOrderUseCase_Handle_Call) RunAndReturn(run func(context.Context, usecase.CreateOrderInput) (usecase.CreateOrderOutput, error)) *MockCreateOrderUseCase_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreateOrderUseCase creates a new instance of MockCreateOrderUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreateOrderUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreateOrderUseCase {
	mock := &MockCreateOrderUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
