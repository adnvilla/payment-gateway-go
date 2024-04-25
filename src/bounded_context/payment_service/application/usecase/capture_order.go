package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
)

type CaptureOrderUseCase interface {
	Handle(ctx context.Context, input CaptureOrderInput) (CaptureOrderOutput, error)
}

type CaptureOrderInput struct {
	OrderId      string
	ProviderType shared_domain.ProviderType
}

type CaptureOrderOutput struct {
	Id        string
	Amount    string
	Currency  string
	CreatedAt int64
}

type captureOrderUseCase struct {
	service         service.OrderProviderService
	orderRepository repository.OrderRepository
}

func NewCaptureOrderUseCase(service service.OrderProviderService, r repository.OrderRepository) use_case.UseCase[CaptureOrderInput, CaptureOrderOutput] {
	u := new(captureOrderUseCase)
	u.service = service
	u.orderRepository = r
	return u
}

func (u *captureOrderUseCase) Handle(ctx context.Context, input CaptureOrderInput) (output CaptureOrderOutput, err error) {
	info := vo.CaptureOrder{
		OrderId:      input.OrderId,
		ProviderType: input.ProviderType,
	}
	output = CaptureOrderOutput{}
	r, err := u.service.CaptureOrder(ctx, info)
	if err != nil {
		return
	}

	err = u.orderRepository.CaptureOrder(ctx, r)
	if err != nil {
		return
	}

	output.Id = r.Id
	output.Amount = r.Amount
	output.Currency = r.Currency
	output.CreatedAt = r.CreatedAt
	return
}
