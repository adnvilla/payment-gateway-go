package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	usecasemock "github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/interfaces/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func getTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func mockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	if content == nil {
		return
	}

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestMockJson(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)
	mockJsonPost(ctx, nil)
}

func TestCreateRefund(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusOK
	response := dto.CreateRefundResponse{}

	mockJsonPost(ctx, body)

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
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.CreateRefundRequest{}
	expected := dto.CreateRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.CreateRefundResponse{}

	mockJsonPost(ctx, body)

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
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.GetRefundRequest{}
	expected := dto.GetRefundResponse{}
	expectedStatus := http.StatusOK
	response := dto.GetRefundResponse{}

	mockJsonPost(ctx, body)

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
	ctx := getTestGinContext(w)

	// Fixture
	body := dto.GetRefundRequest{}
	expected := dto.GetRefundResponse{}
	expectedStatus := http.StatusBadRequest
	response := dto.GetRefundResponse{}

	mockJsonPost(ctx, body)

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
