package provider

import (
	"context"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe/mock"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"go.uber.org/mock/gomock"
)

func TestStripeProvider(t *testing.T) {
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
