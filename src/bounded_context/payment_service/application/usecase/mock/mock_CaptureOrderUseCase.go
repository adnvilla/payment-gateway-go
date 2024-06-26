// Code generated by mockery v2.42.3. DO NOT EDIT.

package mock

import (
	context "context"

	usecase "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	mock "github.com/stretchr/testify/mock"
)

// MockCaptureOrderUseCase is an autogenerated mock type for the CaptureOrderUseCase type
type MockCaptureOrderUseCase struct {
	mock.Mock
}

type MockCaptureOrderUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCaptureOrderUseCase) EXPECT() *MockCaptureOrderUseCase_Expecter {
	return &MockCaptureOrderUseCase_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, input
func (_m *MockCaptureOrderUseCase) Handle(ctx context.Context, input usecase.CaptureOrderInput) (usecase.CaptureOrderOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 usecase.CaptureOrderOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.CaptureOrderInput) (usecase.CaptureOrderOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.CaptureOrderInput) usecase.CaptureOrderOutput); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(usecase.CaptureOrderOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.CaptureOrderInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockCaptureOrderUseCase_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockCaptureOrderUseCase_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - input usecase.CaptureOrderInput
func (_e *MockCaptureOrderUseCase_Expecter) Handle(ctx interface{}, input interface{}) *MockCaptureOrderUseCase_Handle_Call {
	return &MockCaptureOrderUseCase_Handle_Call{Call: _e.mock.On("Handle", ctx, input)}
}

func (_c *MockCaptureOrderUseCase_Handle_Call) Run(run func(ctx context.Context, input usecase.CaptureOrderInput)) *MockCaptureOrderUseCase_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(usecase.CaptureOrderInput))
	})
	return _c
}

func (_c *MockCaptureOrderUseCase_Handle_Call) Return(_a0 usecase.CaptureOrderOutput, _a1 error) *MockCaptureOrderUseCase_Handle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockCaptureOrderUseCase_Handle_Call) RunAndReturn(run func(context.Context, usecase.CaptureOrderInput) (usecase.CaptureOrderOutput, error)) *MockCaptureOrderUseCase_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCaptureOrderUseCase creates a new instance of MockCaptureOrderUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCaptureOrderUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCaptureOrderUseCase {
	mock := &MockCaptureOrderUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
