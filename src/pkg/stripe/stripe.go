package stripe

import (
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/client"
	"github.com/stripe/stripe-go/v78/paymentintent"
	"github.com/stripe/stripe-go/v78/refund"
)

type wrapStripeClient struct {
	c *client.API
}

func NewWrapStripeProvider(key string) *wrapStripeClient {
	sc := GetStripeClient(key)
	return &wrapStripeClient{
		c: sc,
	}
}

func NewWrapStripeProviderWithClients(key string, paymentIntents *paymentintent.Client, refunds *refund.Client) *wrapStripeClient {
	sc := GetStripeClient(key)
	sc.PaymentIntents = paymentIntents
	sc.Refunds = refunds
	return &wrapStripeClient{
		c: sc,
	}
}

func GetStripeClient(key string) *client.API {
	sc := client.New(key, nil) // the second parameter overrides the backends used if needed for mocking
	return sc
}

func (w *wrapStripeClient) NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return w.c.PaymentIntents.New(params)
}

func (w *wrapStripeClient) Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error) {
	return w.c.PaymentIntents.Capture(id, params)
}

func (w *wrapStripeClient) NewRefund(params *stripe.RefundParams) (*stripe.Refund, error) {
	return w.c.Refunds.New(params)
}
