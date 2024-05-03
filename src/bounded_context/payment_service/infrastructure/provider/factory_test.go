package provider

import (
	"context"
	"os"
	"testing"

	paypal_sdk "github.com/adnvilla/payment-gateway-go/src/pkg/paypal"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/stripe"
	"github.com/stretchr/testify/assert"
)

func TestStripeGetProviderClient(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_PROVIDER_STRIPE_KEY", "")
	factory := NewGetProviderFactory()

	expected := NewStripeProvider(stripe.NewWrapStripeProvider(""))
	c, err := factory.GetProviderClient(context.Background(), shared_domain.ProviderType_Stripe)

	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, expected, c)
}

func TestPaypalGetProviderClient(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_PROVIDER_PAYPAL_SECRETID", "qeqw")
	os.Setenv("PAYMENT_GATEWAY_PROVIDER_PAYPAL_CLIENTID", "qwe")
	factory := NewGetProviderFactory()

	expected := NewPaypalProvider(paypal_sdk.GetPaypalClient("qwe", "qeqw"))
	c, err := factory.GetProviderClient(context.Background(), shared_domain.ProviderType_Paypal)

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
