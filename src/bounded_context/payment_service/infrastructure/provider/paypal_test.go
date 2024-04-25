package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	mockProvider "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/provider/mock"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/plutov/paypal/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestPaypalProviderCaptureOrder(t *testing.T) {
	mockPaypal := mockProvider.NewMockPaypalProvider(t)
	ctx := context.Background()
	id := "clientId"

	mockPaypal.EXPECT().CaptureOrder(ctx, id, paypal.CaptureOrderRequest{}).Return(&paypal.CaptureOrderResponse{
		ID: id,
	}, nil).Once()

	paypalProvider := NewPaypalProvider(mockPaypal)

	paypalProvider.CaptureOrder(ctx, vo.CaptureOrder{
		OrderId:      id,
		ProviderType: shared_domain.ProviderType_Paypal,
	})
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
