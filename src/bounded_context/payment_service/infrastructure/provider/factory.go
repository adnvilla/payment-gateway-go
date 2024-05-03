package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	paypal_sdk "github.com/adnvilla/payment-gateway-go/src/pkg/paypal"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe"
)

type factory struct {
}

func NewGetProviderFactory() service.GetProviderService {
	return &factory{}
}

func (f *factory) GetProviderClient(ctx context.Context, provider shared_domain.ProviderType) (service.OrderProviderService, error) {
	switch provider {
	case shared_domain.ProviderType_Stripe:
		key := os.Getenv("PAYMENT_GATEWAY_PROVIDER_STRIPE_KEY")
		client := stripe.NewWrapStripeProvider(key)
		return NewStripeProvider(client), nil
	case shared_domain.ProviderType_Paypal:
		clientid := os.Getenv("PAYMENT_GATEWAY_PROVIDER_PAYPAL_CLIENTID")
		secretid := os.Getenv("PAYMENT_GATEWAY_PROVIDER_PAYPAL_SECRETID")
		client := paypal_sdk.GetPaypalClient(clientid, secretid)
		return NewPaypalProvider(client), nil
	default:
		return nil, fmt.Errorf("error: provider not supported: %v", provider)
	}
}
