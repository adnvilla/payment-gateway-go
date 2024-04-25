package repository

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/entity"
)

type ReundRepository interface {
	CreateRefund(ctx context.Context, refund entity.Refund) error
	GetRefund(ctx context.Context, refund entity.Refund) error
}
