package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	stripewrap "github.com/adnvilla/payment-gateway-go/src/pkg/stripe"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe/mock"
	"github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/refund"
	"go.uber.org/mock/gomock"
)

func TestStripeProviderCreateOrder(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}

	amount := "152"
	currency := "mxn"
	paymentIntent := stripe.PaymentIntent{
		ID:       "int_123",
		Amount:   152,
		Currency: stripe.CurrencyMXN,
	}
	payload, _ := json.Marshal(paymentIntent)

	expected := vo.CreateOrderDetail{
		OrderId:      "int_123",
		Amount:       amount,
		Currency:     currency,
		ProviderType: shared_domain.ProviderType_Stripe,
		Payload:      string(payload),
	}
	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/payment_intents", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(nil).
		Do(func(method string, path string, key string, params stripe.ParamsContainer, v *stripe.PaymentIntent) {
			// Set the return value for the method
			*v = paymentIntent
		}).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	result, err := stripeProvider.CreateOrder(context.Background(), vo.CreateOrder{
		ProviderType: shared_domain.ProviderType_Stripe,
		Amount:       amount,
		Currency:     currency,
	})

	r.NoError(err)
	r.Equal(expected, result)
}

func TestStripeProviderCreateOrderFail(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/payment_intents", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(fmt.Errorf("error")).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	_, err := stripeProvider.CreateOrder(context.Background(), vo.CreateOrder{
		Amount:       "152",
		Currency:     "MXN",
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	r.Error(err)
}

func TestStripeProviderCreateOrderAmountFail(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	_, err := stripeProvider.CreateOrder(context.Background(), vo.CreateOrder{
		Amount:       "152a",
		Currency:     "MXN",
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	r.Error(err)
}

func TestStripeProviderCreateOrderCurrencyFail(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	_, err := stripeProvider.CreateOrder(context.Background(), vo.CreateOrder{
		Amount:       "152",
		Currency:     "COP",
		ProviderType: shared_domain.ProviderType_Stripe,
	})
	r.Error(err)
}

func TestStripeProviderCaptureOrder(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}
	id := "clientId"
	paymentIntent := stripe.PaymentIntent{
		ID: id,
	}
	payload, _ := json.Marshal(paymentIntent)
	expected := vo.CaptureOrderDetail{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Stripe,
		Payload:        string(payload),
	}

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", fmt.Sprintf("/v1/payment_intents/%s/capture", id), gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(nil).
		Do(func(method string, path string, key string, params stripe.ParamsContainer, v *stripe.PaymentIntent) {
			// Set the return value for the method
			*v = paymentIntent
		}).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	result, err := stripeProvider.CaptureOrder(context.Background(), vo.CaptureOrder{
		OrderId:      id,
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	r.NoError(err)
	r.Equal(expected, result)
}

func TestStripeProviderCaptureOrderFail(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &paymentintent.Client{B: mockBackend, Key: "key_123"}
	id := "clientId"

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", fmt.Sprintf("/v1/payment_intents/%s/capture", id), gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(fmt.Errorf("error")).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("", c, nil))

	_, err := stripeProvider.CaptureOrder(context.Background(), vo.CaptureOrder{
		OrderId:      id,
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	r.Error(err)
}

func TestStripeProviderCreateRefund(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &refund.Client{B: mockBackend, Key: "key_123"}
	id := "clientId"

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/refunds", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(nil).
		Do(func(method string, path string, key string, params stripe.ParamsContainer, v *stripe.Refund) {
			// Set the return value for the method
			*v = stripe.Refund{
				ID: id,
			}
		}).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("key_123", nil, c))

	stripeProvider.CreateRefund(context.Background(), vo.CreateRefundOrder{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Stripe,
	})
}

func TestStripeProviderCreateRefundFail(t *testing.T) {
	r := require.New(t)
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	// Create a mock stripe backend
	mockBackend := mock.NewMockBackend(mockController)
	c := &refund.Client{B: mockBackend, Key: "key_123"}
	id := "clientId"

	// Set up a mock call
	mockBackend.EXPECT().Call("POST", "/v1/refunds", gomock.Any(), gomock.Any(), gomock.Any()).
		// Return nil error
		Return(fmt.Errorf("error")).Times(1)

	stripeProvider := NewStripeProvider(stripewrap.NewWrapStripeProviderWithClients("key_123", nil, c))

	_, err := stripeProvider.CreateRefund(context.Background(), vo.CreateRefundOrder{
		CaptureOrderId: id,
		ProviderType:   shared_domain.ProviderType_Stripe,
	})

	r.Error(err)
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
			actualCurrency, err := parseStripeCurrency(tt.currency)
			if !tt.expectedErr && err != nil {
				t.Errorf("Unexpected err but got err: %v", err)
			}
			if !tt.expectedErr && string(actualCurrency) != tt.expected {
				t.Errorf("got %v, want %v", actualCurrency, tt.expected)
			}
		})
	}
}
