package provider

import (
	"context"
	"os"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe"
	"github.com/stretchr/testify/assert"
)

func TestGetGetProviderClient(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_PROVIDER_STRIPE_KEY", "")
	factory := NewGetProviderFactory()

	expected := NewStripeProvider(stripe.GetStripeClient("").PaymentIntents)
	c, err := factory.GetProviderClient(context.Background(), shared_domain.ProviderType_Stripe)

	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, expected, c)
}

func TestGetGetProviderClientNotSupported(t *testing.T) {
	factory := NewGetProviderFactory()

	c, err := factory.GetProviderClient(context.Background(), 100)

	assert.Error(t, err)
	assert.Nil(t, c)
}
