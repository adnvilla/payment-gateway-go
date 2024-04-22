package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	usecasemock "github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/interfaces/dto"
	"github.com/adnvilla/payment-gateway-go/src/pkg/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateRefund(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := testutils.GetTestGinContext(w)

	// Fixture
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusOK
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body)

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
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}

	testutils.MockJsonPost(ctx, body)

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

	testutils.MockJsonPost(ctx, body)

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

	testutils.MockJsonPost(ctx, body)

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
