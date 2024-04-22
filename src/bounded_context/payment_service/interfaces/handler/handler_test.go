package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	usecasemock "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/interfaces/dto"
	"github.com/adnvilla/payment-gateway-go/src/pkg/testutils"
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
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusOK
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body)

	usecaseMock := usecasemock.NewMockCaptureOrderUseCase(t)
	usecaseMock.On("Handle", mock.Anything, mock.Anything).Return(usecase.CaptureOrderOutput{}, nil)

	// Act
	handler := NewCaptureOrderHandler(usecaseMock)
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
	body := dto.CaptureOrderRequest{}
	expected := dto.CaptureOrderResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CaptureOrderResponse{}

	testutils.MockJsonPost(ctx, body)

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
