package service

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
)

type OrderService interface {
	CreateOrder(ctx context.Context, createOrder vo.CreateOrder) (vo.CreateOrderDetail, error)
	CaptureOrder(ctx context.Context, createOrder vo.CaptureOrder) (vo.CaptureOrderDetail, error)
}

type OrderProviderService interface {
	OrderService
}

type createOrderService struct {
	factory GetProviderService
}

func NewCreateOrderService(factory GetProviderService) OrderService {
	c := &createOrderService{
		factory: factory,
	}
	return c
}

func (c *createOrderService) CreateOrder(ctx context.Context, createOrder vo.CreateOrder) (vo.CreateOrderDetail, error) {
	provider, err := c.factory.GetProviderClient(ctx, createOrder.ProviderType)
	if err != nil {
		return vo.CreateOrderDetail{}, err
	}
	return provider.CreateOrder(ctx, createOrder)
}

func (c *createOrderService) CaptureOrder(ctx context.Context, captureOrder vo.CaptureOrder) (vo.CaptureOrderDetail, error) {
	provider, err := c.factory.GetProviderClient(ctx, captureOrder.ProviderType)
	if err != nil {
		return vo.CaptureOrderDetail{}, err
	}
	return provider.CaptureOrder(ctx, captureOrder)
}
