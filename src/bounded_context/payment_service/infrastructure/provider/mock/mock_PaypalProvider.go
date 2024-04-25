// Code generated by mockery v2.42.2. DO NOT EDIT.

package mock

import (
	context "context"

	paypal "github.com/plutov/paypal/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockPaypalProvider is an autogenerated mock type for the PaypalProvider type
type MockPaypalProvider struct {
	mock.Mock
}

type MockPaypalProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaypalProvider) EXPECT() *MockPaypalProvider_Expecter {
	return &MockPaypalProvider_Expecter{mock: &_m.Mock}
}

// CaptureOrder provides a mock function with given fields: ctx, orderID, captureOrderRequest
func (_m *MockPaypalProvider) CaptureOrder(ctx context.Context, orderID string, captureOrderRequest paypal.CaptureOrderRequest) (*paypal.CaptureOrderResponse, error) {
	ret := _m.Called(ctx, orderID, captureOrderRequest)

	if len(ret) == 0 {
		panic("no return value specified for CaptureOrder")
	}

	var r0 *paypal.CaptureOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, paypal.CaptureOrderRequest) (*paypal.CaptureOrderResponse, error)); ok {
		return rf(ctx, orderID, captureOrderRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, paypal.CaptureOrderRequest) *paypal.CaptureOrderResponse); ok {
		r0 = rf(ctx, orderID, captureOrderRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paypal.CaptureOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, paypal.CaptureOrderRequest) error); ok {
		r1 = rf(ctx, orderID, captureOrderRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaypalProvider_CaptureOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CaptureOrder'
type MockPaypalProvider_CaptureOrder_Call struct {
	*mock.Call
}

// CaptureOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - orderID string
//   - captureOrderRequest paypal.CaptureOrderRequest
func (_e *MockPaypalProvider_Expecter) CaptureOrder(ctx interface{}, orderID interface{}, captureOrderRequest interface{}) *MockPaypalProvider_CaptureOrder_Call {
	return &MockPaypalProvider_CaptureOrder_Call{Call: _e.mock.On("CaptureOrder", ctx, orderID, captureOrderRequest)}
}

func (_c *MockPaypalProvider_CaptureOrder_Call) Run(run func(ctx context.Context, orderID string, captureOrderRequest paypal.CaptureOrderRequest)) *MockPaypalProvider_CaptureOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(paypal.CaptureOrderRequest))
	})
	return _c
}

func (_c *MockPaypalProvider_CaptureOrder_Call) Return(_a0 *paypal.CaptureOrderResponse, _a1 error) *MockPaypalProvider_CaptureOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaypalProvider_CaptureOrder_Call) RunAndReturn(run func(context.Context, string, paypal.CaptureOrderRequest) (*paypal.CaptureOrderResponse, error)) *MockPaypalProvider_CaptureOrder_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOrder provides a mock function with given fields: ctx, intent, purchaseUnits, payer, appContext
func (_m *MockPaypalProvider) CreateOrder(ctx context.Context, intent string, purchaseUnits []paypal.PurchaseUnitRequest, payer *paypal.CreateOrderPayer, appContext *paypal.ApplicationContext) (*paypal.Order, error) {
	ret := _m.Called(ctx, intent, purchaseUnits, payer, appContext)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 *paypal.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, []paypal.PurchaseUnitRequest, *paypal.CreateOrderPayer, *paypal.ApplicationContext) (*paypal.Order, error)); ok {
		return rf(ctx, intent, purchaseUnits, payer, appContext)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, []paypal.PurchaseUnitRequest, *paypal.CreateOrderPayer, *paypal.ApplicationContext) *paypal.Order); ok {
		r0 = rf(ctx, intent, purchaseUnits, payer, appContext)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*paypal.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, []paypal.PurchaseUnitRequest, *paypal.CreateOrderPayer, *paypal.ApplicationContext) error); ok {
		r1 = rf(ctx, intent, purchaseUnits, payer, appContext)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPaypalProvider_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type MockPaypalProvider_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - intent string
//   - purchaseUnits []paypal.PurchaseUnitRequest
//   - payer *paypal.CreateOrderPayer
//   - appContext *paypal.ApplicationContext
func (_e *MockPaypalProvider_Expecter) CreateOrder(ctx interface{}, intent interface{}, purchaseUnits interface{}, payer interface{}, appContext interface{}) *MockPaypalProvider_CreateOrder_Call {
	return &MockPaypalProvider_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx, intent, purchaseUnits, payer, appContext)}
}

func (_c *MockPaypalProvider_CreateOrder_Call) Run(run func(ctx context.Context, intent string, purchaseUnits []paypal.PurchaseUnitRequest, payer *paypal.CreateOrderPayer, appContext *paypal.ApplicationContext)) *MockPaypalProvider_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].([]paypal.PurchaseUnitRequest), args[3].(*paypal.CreateOrderPayer), args[4].(*paypal.ApplicationContext))
	})
	return _c
}

func (_c *MockPaypalProvider_CreateOrder_Call) Return(_a0 *paypal.Order, _a1 error) *MockPaypalProvider_CreateOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaypalProvider_CreateOrder_Call) RunAndReturn(run func(context.Context, string, []paypal.PurchaseUnitRequest, *paypal.CreateOrderPayer, *paypal.ApplicationContext) (*paypal.Order, error)) *MockPaypalProvider_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPaypalProvider creates a new instance of MockPaypalProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPaypalProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPaypalProvider {
	mock := &MockPaypalProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
