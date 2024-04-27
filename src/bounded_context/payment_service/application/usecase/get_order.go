package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	uuid "github.com/satori/go.uuid"
)

type GetOrderUseCase interface {
	Handle(ctx context.Context, input GetOrderInput) (GetOrderOutput, error)
}

type GetOrderInput struct {
	Id uuid.UUID
}

type GetOrderOutput struct {
	Id           uuid.UUID
	Amount       string
	Currency     string
	CreatedAt    int64
	ProviderType shared_domain.ProviderType
}

type getOrderUseCase struct {
	orderRepository repository.OrderRepository
}

func NewGetOrderUseCase(r repository.OrderRepository) use_case.UseCase[GetOrderInput, GetOrderOutput] {
	u := new(getOrderUseCase)
	u.orderRepository = r
	return u
}

func (u *getOrderUseCase) Handle(ctx context.Context, input GetOrderInput) (output GetOrderOutput, err error) {
	output = GetOrderOutput{}

	order, err := u.orderRepository.GetOrder(ctx, input.Id)
	if err != nil {
		return
	}

	output.Id = order.Id
	output.Amount = order.Amount
	output.Currency = order.Currency
	output.CreatedAt = order.CreatedAt
	output.ProviderType = order.ProviderType
	return output, nil
}
