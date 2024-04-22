package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
)

type CreateRefundUseCase interface {
	Handle(ctx context.Context, input CreateRefundInput) (CreateRefundOutput, error)
}

type CreateRefundInput struct{}

type CreateRefundOutput struct{}

type createRefundUseCase struct{}

func NewCreateRefundUseCase() use_case.UseCase[CreateRefundInput, CreateRefundOutput] {
	u := new(createRefundUseCase)
	return u
}

func (u *createRefundUseCase) Handle(ctx context.Context, input CreateRefundInput) (CreateRefundOutput, error) {
	output := CreateRefundOutput{}
	return output, nil
}
