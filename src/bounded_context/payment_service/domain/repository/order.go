package repository

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order vo.CreateOrderDetail) error
	CaptureOrder(ctx context.Context, order vo.CaptureOrderDetail) error
}
