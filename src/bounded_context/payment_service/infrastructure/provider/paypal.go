package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/plutov/paypal/v4"
)

type PaypalProvider interface {
	CreateOrder(ctx context.Context, intent string, purchaseUnits []paypal.PurchaseUnitRequest, payer *paypal.CreateOrderPayer, appContext *paypal.ApplicationContext) (*paypal.Order, error)
	CaptureOrder(ctx context.Context, orderID string, captureOrderRequest paypal.CaptureOrderRequest) (*paypal.CaptureOrderResponse, error)
}

type paypalProvider struct {
	paypalClient PaypalProvider
}

func NewPaypalProvider(stripProvider PaypalProvider) service.OrderProviderService {
	return &paypalProvider{
		paypalClient: stripProvider,
	}
}

func (s *paypalProvider) CreateOrder(ctx context.Context, createOrder vo.CreateOrder) (vo.CreateOrderDetail, error) {
	amount, err := strconv.ParseInt(createOrder.Amount, 10, 64)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	currency, err := parsePaypalCurrency(createOrder.Currency)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	result, err := s.paypalClient.CreateOrder(ctx, paypal.OrderIntentCapture, []paypal.PurchaseUnitRequest{
		{
			Amount: &paypal.PurchaseUnitAmount{
				Currency: string(currency),
				Value:    strconv.FormatInt(amount, 10),
			},
		},
	}, nil, nil)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}

	return vo.CreateOrderDetail{
		OrderId:      result.ID,
		Amount:       strconv.FormatInt(amount, 10),
		Currency:     string(currency),
		ProviderType: createOrder.ProviderType,
		Payload:      string(payload),
	}, nil
}

func (s *paypalProvider) CaptureOrder(ctx context.Context, captureOrder vo.CaptureOrder) (vo.CaptureOrderDetail, error) {

	result, err := s.paypalClient.CaptureOrder(ctx, captureOrder.OrderId, paypal.CaptureOrderRequest{})

	if err != nil {
		return vo.CaptureOrderDetail{}, err
	}

	payload, err := json.Marshal(result)
	if err != nil {
		return vo.CaptureOrderDetail{}, err
	}

	return vo.CaptureOrderDetail{
		CaptureOrderId: result.ID,
		ProviderType:   captureOrder.ProviderType,
		Payload:        string(payload),
	}, nil
}

func parsePaypalCurrency(c string) (string, error) {
	switch c {
	case "USD", "usd":
		return "USD", nil
	case "MXN", "mxn":
		return "MXN", nil
	default:
		return "", fmt.Errorf("currency not supported %v", c)
	}
}
