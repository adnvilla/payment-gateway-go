package provider

import (
	"context"
	"fmt"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
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
		client := stripe.GetStripeClient("sk_test")
		return NewStripeProvider(client.PaymentIntents), nil
	case shared_domain.ProviderType_Paypal:
		return nil, fmt.Errorf("error: provider not supported: %v", provider)
	default:
		return nil, fmt.Errorf("error: provider not supported: %v", provider)
	}
}
