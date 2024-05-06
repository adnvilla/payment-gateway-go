package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	mockProvider "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/provider/mock"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/plutov/paypal/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestPaypalProviderCreateOrder(t *testing.T) {
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"
	currency := "MXN"
	amount := "152"

	mockPaypal.EXPECT().CreateOrder(ctx, paypal.OrderIntentCapture, mock.Anything, mock.Anything, mock.Anything).Return(&paypal.Order{
		ID: id,
	}, nil).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	r, err := paypalProvider.CreateOrder(ctx, vo.CreateOrder{
		Amount:       amount,
		Currency:     currency,
		ProviderType: shared_domain.ProviderType_Paypal,
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestPaypalProviderCreateOrderFail(t *testing.T) {
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	currency := "MXN"
	amount := "152"

	mockPaypal.EXPECT().CreateOrder(ctx, paypal.OrderIntentCapture, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error")).Once()
	expected := vo.CreateOrderDetail{}

	paypalProvider := NewPaypalProvider(mockPaypal)

	r, err := paypalProvider.CreateOrder(ctx, vo.CreateOrder{
		Amount:       amount,
		Currency:     currency,
		ProviderType: shared_domain.ProviderType_Paypal,
	})

	assert.Error(t, err)
	assert.Equal(t, expected, r)
}

func TestPaypalProviderCreateOrderCurrencyFail(t *testing.T) {
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	currency := "COP"
	amount := "152"

	paypalProvider := NewPaypalProvider(mockPaypal)

	r, err := paypalProvider.CreateOrder(ctx, vo.CreateOrder{
		Amount:       amount,
		Currency:     currency,
		ProviderType: shared_domain.ProviderType_Paypal,
	})

	assert.Error(t, err)
	assert.Equal(t, vo.CreateOrderDetail{}, r)
}

func TestPaypalProviderCreateOrderAmountFail(t *testing.T) {
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	currency := "MXN"

	paypalProvider := NewPaypalProvider(mockPaypal)

	r, err := paypalProvider.CreateOrder(ctx, vo.CreateOrder{
		Amount:       "error",
		Currency:     currency,
		ProviderType: shared_domain.ProviderType_Paypal,
	})

	assert.Error(t, err)
	assert.Equal(t, vo.CreateOrderDetail{}, r)
}

func TestPaypalProviderCaptureOrder(t *testing.T) {
	r := require.New(t)
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"

	captureResponse := &paypal.CaptureOrderResponse{
		ID: id,
	}
	payload, _ := json.Marshal(captureResponse)
	expected := vo.CaptureOrderDetail{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Paypal,
		Payload:        string(payload),
	}

	mockPaypal.EXPECT().CaptureOrder(ctx, id, paypal.CaptureOrderRequest{}).Return(captureResponse, nil).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	reult, err := paypalProvider.CaptureOrder(ctx, vo.CaptureOrder{
		OrderId:      id,
		ProviderType: shared_domain.ProviderType_Paypal,
	})

	r.NoError(err)
	r.Equal(expected, reult)
}

func TestPaypalProviderCaptureOrderFail(t *testing.T) {
	r := require.New(t)
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"

	mockPaypal.EXPECT().CaptureOrder(ctx, id, paypal.CaptureOrderRequest{}).Return(nil, fmt.Errorf("error")).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	_, err := paypalProvider.CaptureOrder(ctx, vo.CaptureOrder{
		OrderId:      id,
		ProviderType: shared_domain.ProviderType_Paypal,
	})
	r.Error(err)
}

func TestPaypalProviderCreateRefund(t *testing.T) {
	r := require.New(t)
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"
	refundResponse := &paypal.RefundResponse{
		ID: id,
	}
	payload, _ := json.Marshal(refundResponse)
	expected := vo.CreateRefundDetail{
		RefundOrderId: id,
		ProviderType:  shared_domain.ProviderType_Paypal,
		Payload:       string(payload),
	}

	mockPaypal.EXPECT().RefundCapture(ctx, id, paypal.RefundCaptureRequest{}).Return(refundResponse, nil).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	result, err := paypalProvider.CreateRefund(ctx, vo.CreateRefundOrder{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Paypal,
	})

	r.NoError(err)
	r.Equal(expected, result)
}

func TestPaypalProviderCreateRefundFail(t *testing.T) {
	r := require.New(t)
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"

	mockPaypal.EXPECT().RefundCapture(ctx, id, paypal.RefundCaptureRequest{}).Return(nil, fmt.Errorf("error")).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	_, err := paypalProvider.CreateRefund(ctx, vo.CreateRefundOrder{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Paypal,
	})

	r.Error(err)
}

func TestParsePaypalCurrency(t *testing.T) {
	var tests = []struct {
		currency    string
		expected    string
		expectedErr bool
	}{
		{"MXN", "MXN", false},
		{"mxn", "MXN", false},
		{"USD", "USD", false},
		{"usd", "USD", false},
		{"COP", "COP", true},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%v,%v", tt.currency, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actualCurrency, err := parsePaypalCurrency(tt.currency)
			if !tt.expectedErr && err != nil {
				t.Errorf("Unexpected err but got err: %v", err)
			}
			if !tt.expectedErr && string(actualCurrency) != tt.expected {
				t.Errorf("got %v, want %v", actualCurrency, tt.expected)
			}
		})
	}

}
