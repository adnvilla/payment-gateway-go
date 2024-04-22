package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
)

type GetRefundUseCase interface {
	Handle(ctx context.Context, input GetRefundInput) (GetRefundOutput, error)
}

type GetRefundInput struct{}

type GetRefundOutput struct{}

type getRefundUseCase struct{}

func NewGetRefundUseCase() use_case.UseCase[GetRefundInput, GetRefundOutput] {
	u := new(getRefundUseCase)
	return u
}

func (u *getRefundUseCase) Handle(ctx context.Context, input GetRefundInput) (GetRefundOutput, error) {
	output := GetRefundOutput{}
	return output, nil
}
