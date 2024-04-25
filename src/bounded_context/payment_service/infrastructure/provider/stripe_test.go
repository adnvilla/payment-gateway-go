package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe/mock"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"go.uber.org/mock/gomock"
)

func TestStripeProviderCreateOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := paymentintent.Client{B: mockBackend, Key: "key_123"}

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/payment_intents", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(nil).
		Do(func(method string, path string, key string, params stripe.ParamsContainer, v *stripe.PaymentIntent) {
			// Set the return value for the method
			*v = stripe.PaymentIntent{
				ID: "int_123",
			}
		}).Times(1)

	stripeProvider := NewStripeProvider(c)

	stripeProvider.CreateOrder(context.Background(), vo.CreateOrder{
		Amount:       "152",
		Currency:     "MXN",
		ProviderType: shared_domain.ProviderType_Stripe,
	})
}

func TestStripeProviderCaptureOrder(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := paymentintent.Client{B: mockBackend, Key: "key_123"}

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/payment_intents/int_123/capture", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(nil).
		Do(func(method string, path string, key string, params stripe.ParamsContainer, v *stripe.PaymentIntent) {
			// Set the return value for the method
			*v = stripe.PaymentIntent{
				ID: "int_123",
			}
		}).Times(1)

	stripeProvider := NewStripeProvider(c)

	stripeProvider.CaptureOrder(context.Background(), vo.CaptureOrder{
		OrderId:      "int_123",
		ProviderType: shared_domain.ProviderType_Stripe,
	})
}

func TestParseCurrency(t *testing.T) {
	var tests = []struct {
		currency    string
		expected    string
		expectedErr bool
	}{
		{"MXN", "mxn", false},
		{"mxn", "mxn", false},
		{"USD", "usd", false},
		{"usd", "usd", false},
		{"COP", "cop", true},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("%v,%v", tt.currency, tt.expected)
		t.Run(testname, func(t *testing.T) {
			actualCurrency, err := parseCurrency(tt.currency)
			if !tt.expectedErr && err != nil {
				t.Errorf("Unexpected err but got err: %v", err)
			}
			if !tt.expectedErr && string(actualCurrency) != tt.expected {
				t.Errorf("got %v, want %v", tt.currency, tt.expected)
			}
		})
	}

}
