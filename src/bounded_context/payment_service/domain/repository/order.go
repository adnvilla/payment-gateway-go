package repository

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	uuid "github.com/satori/go.uuid"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order vo.CreateOrderDetail) (uuid.UUID, error)
	GetOrder(ctx context.Context, orderId uuid.UUID) (vo.CreateOrder, error)
	GetOrderProvider(ctx context.Context, orderId uuid.UUID) (vo.CreateOrderDetail, error)
	CaptureOrder(ctx context.Context, order vo.CaptureOrderDetail) (uuid.UUID, error)
}
