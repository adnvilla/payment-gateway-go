package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
)

type CreateOrderUseCase interface {
	Handle(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error)
}

type CreateOrderInput struct{}

type CreateOrderOutput struct{}

type createOrderUseCase struct{}

func NewCreateOrderUseCase() use_case.UseCase[CreateOrderInput, CreateOrderOutput] {
	u := new(createOrderUseCase)
	return u
}

func (u *createOrderUseCase) Handle(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error) {
	output := CreateOrderOutput{}
	return output, nil
}
