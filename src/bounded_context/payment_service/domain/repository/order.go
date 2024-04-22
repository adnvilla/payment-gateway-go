package repository

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/entity"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order entity.Order) error
	CaptureOrder(ctx context.Context, order entity.Order) error
}
