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

type CreateRefundUseCase interface {
	Handle(ctx context.Context, input CreateRefundInput) (CreateRefundOutput, error)
}

type CreateRefundInput struct {
	CaptureOrderId uuid.UUID
	ProviderType   shared_domain.ProviderType
}

type CreateRefundOutput struct {
	Id        uuid.UUID
	Amount    string
	Currency  string
	CreatedAt int64
}

type createRefundUseCase struct {
	service         service.OrderProviderService
	orderRepository repository.OrderRepository
}

func NewCreateRefundUseCase(service service.OrderProviderService, r repository.OrderRepository) use_case.UseCase[CreateRefundInput, CreateRefundOutput] {
	u := new(createRefundUseCase)
	u.service = service
	u.orderRepository = r
	return u
}

func (u *createRefundUseCase) Handle(ctx context.Context, input CreateRefundInput) (output CreateRefundOutput, err error) {
	output = CreateRefundOutput{}

	order, err := u.orderRepository.GetCaptureOrderProvider(ctx, input.CaptureOrderId)
	if err != nil {
		return
	}

	info := vo.CreateRefundOrder{
		CaptureOrderId: order.CaptureOrderId,
		ProviderType:   order.ProviderType,
	}
	r, err := u.service.CreateRefund(ctx, info)
	if err != nil {
		return
	}

	id, err := u.orderRepository.CreateRefund(ctx, r)
	if err != nil {
		return
	}

	output.Id = id
	//output.Amount = r.Amount
	//output.Currency = r.Currency
	//output.CreatedAt = r.CreatedAt
	return output, nil
}
