package postgresql

import (
	"context"
	"errors"

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
	orderModel := models.OrderModel{}
	result := r.db.Create(&orderModel)
	if result.Error != nil {
		return errors.New("have a issue with consult DB")
	}

	return nil
}
func (r *orderRepository) CaptureOrder(ctx context.Context, order vo.CaptureOrderDetail) error {

	return nil
}
