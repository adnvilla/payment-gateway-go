package postgresql

import (
	"context"
	"fmt"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/sql/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order vo.CreateOrderDetail) error {
	orderModel := models.CreateOrder{
		Amount:       order.Amount,
		Currency:     order.Currency,
		ProviderType: int(order.ProviderType),
		CreateOrderProvider: models.CreateOrderProvider{
			ProviderType: int(order.ProviderType),
			Payload:      order.Payload,
		},
	}
	result := r.db.Create(&orderModel)
	if result.Error != nil {
		return fmt.Errorf("have a issue with consult DB:", result.Error)
	}

	return nil
}
func (r *orderRepository) CaptureOrder(ctx context.Context, order vo.CaptureOrderDetail) error {

	return nil
}
