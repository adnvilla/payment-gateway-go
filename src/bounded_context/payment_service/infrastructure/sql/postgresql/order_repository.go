package postgresql

import (
	"context"
	"fmt"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/sql/models"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
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

func (r *orderRepository) CreateOrder(ctx context.Context, order vo.CreateOrderDetail) (uuid.UUID, error) {
	orderModel := models.CreateOrder{
		Amount:       order.Amount,
		Currency:     order.Currency,
		ProviderType: int(order.ProviderType),
		CreateOrderProvider: models.CreateOrderProvider{
			ProviderOrderID: order.OrderId,
			ProviderType:    int(order.ProviderType),
			Payload:         order.Payload,
		},
	}
	result := r.db.Create(&orderModel)
	if result.Error != nil {
		return uuid.UUID{}, fmt.Errorf("have a issue with insert DB CreateOrder: %v", result.Error)
	}

	return orderModel.ID, nil
}

func (r *orderRepository) CaptureOrder(ctx context.Context, order vo.CaptureOrderDetail) (uuid.UUID, error) {
	captureModel := models.CaptureOrder{
		ProviderType: int(order.ProviderType),
		CaptureOrderProvider: models.CaptureOrderProvider{
			ProviderCaptureID: order.CaptureOrderId,
			ProviderType:      int(order.ProviderType),
			Payload:           order.Payload,
		},
	}
	result := r.db.Create(&captureModel)
	if result.Error != nil {
		return uuid.UUID{}, fmt.Errorf("have a issue with insert DB CaptureOrder: %v", result.Error)
	}

	return captureModel.ID, nil
}

func (r *orderRepository) GetOrderProvider(ctx context.Context, order uuid.UUID) (vo.CreateOrderDetail, error) {

	orderProvider := models.CreateOrderProvider{
		CreateOrderID: order,
	}

	result := r.db.Where("create_order_id = ?", order.String()).First(&orderProvider)
	if result.Error != nil {
		return vo.CreateOrderDetail{}, fmt.Errorf("have a issue with consult DB CreateOrderProvider: %v", result.Error)
	}

	return vo.CreateOrderDetail{
		Id:           orderProvider.ID,
		OrderId:      orderProvider.ProviderOrderID,
		ProviderType: shared_domain.ProviderType(orderProvider.ProviderType),
	}, nil
}

func (r *orderRepository) GetOrder(ctx context.Context, order uuid.UUID) (vo.CreateOrder, error) {

	orderProvider := models.CreateOrder{}
	orderProvider.ID = order

	result := r.db.First(&orderProvider)
	if result.Error != nil {
		return vo.CreateOrder{}, fmt.Errorf("have a issue with consult DB CreateOrderProvider: %v", result.Error)
	}

	return vo.CreateOrder{
		Id:           orderProvider.ID,
		Amount:       orderProvider.Amount,
		Currency:     orderProvider.Currency,
		CreatedAt:    int64(orderProvider.CreatedAt),
		ProviderType: shared_domain.ProviderType(orderProvider.ProviderType),
	}, nil
}

func (r *orderRepository) CreateRefund(ctx context.Context, order vo.CreateRefundDetail) (uuid.UUID, error) {
	captureModel := models.Refund{
		ProviderType: int(order.ProviderType),
		RefundProvider: models.RefundProvider{
			ProviderRefundID: order.RefundOrderId,
			ProviderType:     int(order.ProviderType),
			Payload:          order.Payload,
		},
	}
	result := r.db.Create(&captureModel)
	if result.Error != nil {
		return uuid.UUID{}, fmt.Errorf("have a issue with insert DB CreateRefund: %v", result.Error)
	}

	return captureModel.ID, nil
}

func (r *orderRepository) GetCaptureOrderProvider(ctx context.Context, order uuid.UUID) (vo.CaptureOrderDetail, error) {

	orderProvider := models.CaptureOrderProvider{
		CaptureOrderID: order,
	}

	result := r.db.Where("capture_order_id = ?", order.String()).First(&orderProvider)
	if result.Error != nil {
		return vo.CaptureOrderDetail{}, fmt.Errorf("have a issue with consult DB CreateOrderProvider: %v", result.Error)
	}

	return vo.CaptureOrderDetail{
		Id:             orderProvider.ID,
		CaptureOrderId: orderProvider.ProviderCaptureID,
		ProviderType:   shared_domain.ProviderType(orderProvider.ProviderType),
	}, nil
}
