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

	testutils.MockJsonPost(ctx, body)

	usecaseMock := usecasemock.NewMockCreateOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateOrderOutput{}, nil)

	// Act
	handler := NewCreateOrderHandler(usecaseMock)
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

	testutils.MockJsonPost(ctx, body)

	// Act
	handler := NewCreateOrderHandler(nil)
	handler.CreateOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateOrderFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateOrderRequest{}
	expected := dto.CreateOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateOrderResponse{}

	testutils.MockJsonPost(ctx, body)

	usecaseMock := usecasemock.NewMockCreateOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateOrderOutput{}, fmt.Errorf("some error"))

	// Act
	handler := NewCreateOrderHandler(usecaseMock)
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

	testutils.MockJsonPost(ctx, body)
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

	// Act
	handler := NewCaptureOrderHandler(usecaseMock)
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

	testutils.MockJsonPost(ctx, body)
	ctx.Params = []gin.Param{
		{
			Key:   "failId",
			Value: "order",
		},
	}

	// Act
	handler := NewCaptureOrderHandler(nil)
	handler.CaptureOrder(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCaptureOrderFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCaptureOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CaptureOrderOutput{}, fmt.Errorf("some error"))

	// Act
	handler := NewCaptureOrderHandler(usecaseMock)
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

	testutils.MockJsonPost(ctx, body)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCreateRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateRefundOutput{}, nil)

	// Act
	handler := NewCreateRefundHandler(usecaseMock)
	handler.CreateRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestCreateRefundFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	id := uuid.NewV4()
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: id.String(),
		},
	}

	usecaseMock := usecasemock.NewMockCreateRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CreateRefundOutput{}, fmt.Errorf("some error"))

	// Act
	handler := NewCreateRefundHandler(usecaseMock)
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
	params := url.Values{}
	params.Add("charge_id", body.Charge)

	testutils.MockJsonGet(ctx, params)

	usecaseMock := usecasemock.NewMockGetRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetRefundOutput{}, nil)

	// Act
	handler := NewGetRefundHandler(usecaseMock)
	handler.GetRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}

func TestGetRefundFail(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.GetRefundRequest{}
	expected := dto.GetRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetRefundResponse{}
	params := url.Values{}
	params.Add("charge_id", body.Charge)

	testutils.MockJsonGet(ctx, params)

	usecaseMock := usecasemock.NewMockGetRefundUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.GetRefundOutput{}, fmt.Errorf("some error"))

	// Act
	handler := NewGetRefundHandler(usecaseMock)
	handler.GetRefund(ctx)

	// Assert
	assert.EqualValues(t, expectedStatus, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expected, response)
}
