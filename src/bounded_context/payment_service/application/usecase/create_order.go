package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	uuid "github.com/satori/go.uuid"
)

type CreateOrderUseCase interface {
	Handle(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error)
}

type CreateOrderInput struct {
	Amount       string
	Currency     string
	ProviderType shared_domain.ProviderType
}

type CreateOrderOutput struct {
	Id        uuid.UUID
	Amount    string
	Currency  string
	CreatedAt int64
}

type createOrderUseCase struct {
	service         service.OrderProviderService
	orderRepository repository.OrderRepository
}

func NewCreateOrderUseCase(service service.OrderProviderService, r repository.OrderRepository) use_case.UseCase[CreateOrderInput, CreateOrderOutput] {
	u := new(createOrderUseCase)
	u.service = service
	u.orderRepository = r
	return u
}

func (u *createOrderUseCase) Handle(ctx context.Context, input CreateOrderInput) (output CreateOrderOutput, err error) {
	output = CreateOrderOutput{}
	info := vo.CreateOrder{
		Amount:       input.Amount,
		Currency:     input.Currency,
		ProviderType: input.ProviderType,
	}
	r, err := u.service.CreateOrder(ctx, info)
	if err != nil {
		return
	}

	id, err := u.orderRepository.CreateOrder(ctx, r)
	if err != nil {
		return
	}

	output.Id = id
	output.Amount = r.Amount
	output.Currency = r.Currency
	output.CreatedAt = r.CreatedAt
	return output, nil
}
