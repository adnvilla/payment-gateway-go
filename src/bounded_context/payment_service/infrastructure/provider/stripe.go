package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/stripe/stripe-go/v78"
)

type StripeProvider interface {
	NewPaymentIntent(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
	Capture(id string, params *stripe.PaymentIntentCaptureParams) (*stripe.PaymentIntent, error)
	NewRefund(params *stripe.RefundParams) (*stripe.Refund, error)
}

type stripeProvider struct {
	stripeClient StripeProvider
}

func NewStripeProvider(stripProvider StripeProvider) service.OrderProviderService {
	return &stripeProvider{
		stripeClient: stripProvider,
	}
}

func (s *stripeProvider) CreateOrder(ctx context.Context, createOrder vo.CreateOrder) (vo.CreateOrderDetail, error) {
	amount, err := strconv.ParseInt(createOrder.Amount, 10, 64)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	currency, err := parseStripeCurrency(createOrder.Currency)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount),
		Currency: stripe.String(string(currency)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}
	result, err := s.stripeClient.NewPaymentIntent(params)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	return vo.CreateOrderDetail{
		OrderId:      result.ID,
		Amount:       strconv.FormatInt(result.Amount, 10),
		CreatedAt:    result.Created,
		Currency:     string(result.Currency),
		ProviderType: createOrder.ProviderType,
		Payload:      string(payload),
	}, nil
}

func (s *stripeProvider) CaptureOrder(ctx context.Context, captureOrder vo.CaptureOrder) (vo.CaptureOrderDetail, error) {
	params := &stripe.PaymentIntentCaptureParams{}
	result, err := s.stripeClient.Capture(captureOrder.OrderId, params)
	if err != nil {
		return vo.CaptureOrderDetail{}, err
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return vo.CaptureOrderDetail{}, err
	}

	return vo.CaptureOrderDetail{
		CaptureOrderId: result.ID,
		Payload:        string(payload),
		ProviderType:   captureOrder.ProviderType,
	}, nil
}

func (s *stripeProvider) CreateRefund(ctx context.Context, captureOrder vo.CreateRefundOrder) (vo.CreateRefundDetail, error) {
	params := &stripe.RefundParams{
		Charge: stripe.String(captureOrder.CaptureOrderId),
	}

	result, err := s.stripeClient.NewRefund(params)
	if err != nil {
		return vo.CreateRefundDetail{}, err
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return vo.CreateRefundDetail{}, err
	}

	return vo.CreateRefundDetail{
		RefundOrderId: result.ID,
		Payload:       string(payload),
	}, nil
}

func parseStripeCurrency(c string) (stripe.Currency, error) {
	switch c {
	case "USD", "usd":
		return stripe.CurrencyUSD, nil
	case "MXN", "mxn":
		return stripe.CurrencyMXN, nil
	default:
		return "", fmt.Errorf("currency not supported %v", c)
	}
}
