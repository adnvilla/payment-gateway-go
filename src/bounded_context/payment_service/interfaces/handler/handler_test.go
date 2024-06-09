package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	usecasemock "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/interfaces/dto"
	"github.com/adnvilla/payment-gateway-go/src/pkg/dispatcher"
	"github.com/adnvilla/payment-gateway-go/src/pkg/testutils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateOrderRequest{}
	expected := dto.CreateOrderResponse{}
	expectedStatus := http.StatusOK
	response := dto.CreateOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)

	usecaseMock := usecasemock.NewMockCreateOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateOrderOutput{}, nil)

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCreateOrderHandler()
	handler.CreateOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateOrderFailBody(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := "{ shd }"
	expected := dto.CreateOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)

	// Act
	handler := NewCreateOrderHandler()
	handler.CreateOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateOrderHandleFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateOrderRequest{}
	expected := dto.CreateOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)

	usecaseMock := usecasemock.NewMockCreateOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateOrderOutput{}, fmt.Errorf("some error"))

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCreateOrderHandler()
	handler.CreateOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCaptureOrder(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{
		Id: id.String(),
	}
	expectedStatus := http.StatusOK
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCaptureOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CaptureOrderOutput{
		Id: id,
	}, nil)

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCaptureOrderHandler()
	handler.CaptureOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCaptureOrderFailParameter(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)
	ctx.Params = []gin.Param{
		{
			Key:   "failId",
			Value: "order",
		},
	}

	// Act
	handler := NewCaptureOrderHandler()
	handler.CaptureOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCaptureOrderHanldeFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCaptureOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CaptureOrderOutput{}, fmt.Errorf("some error"))

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCaptureOrderHandler()
	handler.CaptureOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCaptureOrderFailParameterId(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CaptureOrderResponse{}
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: "failId",
		},
	}

	testutils.MockJsonPost(ctx, body, nil, nil)

	// Act
	handler := NewCaptureOrderHandler()
	handler.CaptureOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateRefund(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusOK
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCreateRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateRefundOutput{}, nil)

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCreateRefundHandler()
	handler.CreateRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateRefundHandleFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCreateRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateRefundOutput{}, fmt.Errorf("some error"))

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewCreateRefundHandler()
	handler.CreateRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateRefundFailParameterId(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: "failId",
		},
	}

	testutils.MockJsonPost(ctx, body, nil, nil)

	// Act
	handler := NewCreateRefundHandler()
	handler.CreateRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateRefundFailParameter(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body, nil, nil)

	// Act
	handler := NewCreateRefundHandler()
	handler.CreateRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetRefund(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.GetRefundRequest{}
	expected := dto.GetRefundResponse{}
	expectedStatus := http.StatusOK
	response := dto.GetRefundResponse{}
	urlValues := url.Values{}
	urlValues.Add("charge_id", body.Charge)

	testutils.MockJsonGet(ctx, nil, urlValues)

	usecaseMock := usecasemock.NewMockGetRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetRefundOutput{}, nil)

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewGetRefundHandler()
	handler.GetRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetRefundHanldeFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.GetRefundRequest{}
	expected := dto.GetRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetRefundResponse{}
	urlValues := url.Values{}
	urlValues.Add("charge_id", body.Charge)

	testutils.MockJsonGet(ctx, nil, urlValues)

	usecaseMock := usecasemock.NewMockGetRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetRefundOutput{}, fmt.Errorf("some error"))

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewGetRefundHandler()
	handler.GetRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetOrder(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	expected := dto.GetOrderResponse{
		Id: id.String(),
	}
	expectedStatus := http.StatusOK
	response := dto.GetOrderResponse{}
	params := []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	testutils.MockJsonGet(ctx, params, nil)

	usecaseMock := usecasemock.NewMockGetOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetOrderOutput{
		Id: id,
	}, nil)

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewGetOrderHandler()
	handler.GetOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetOrderHanldeFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	expected := dto.GetOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetOrderResponse{}
	params := []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	testutils.MockJsonGet(ctx, params, nil)

	usecaseMock := usecasemock.NewMockGetOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetOrderOutput{}, fmt.Errorf("some error"))

	dispatcher.Reset()
	dispatcher.RegisterHandler(usecaseMock)

	// Act
	handler := NewGetOrderHandler()
	handler.GetOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetOrderFailParameter(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	expected := dto.GetOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetOrderResponse{}
	params := []gin.Param{
		{
			Key:   "failId",
			Value: "order",
		},
	}

	testutils.MockJsonGet(ctx, params, nil)

	// Act
	handler := NewGetOrderHandler()
	handler.GetOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetOrderFailParameterId(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	expected := dto.GetOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetOrderResponse{}
	params := []gin.Param{
		{
			Key:   "id",
			Value: "failId",
		},
	}

	testutils.MockJsonGet(ctx, params, nil)

	// Act
	handler := NewGetOrderHandler()
	handler.GetOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}
