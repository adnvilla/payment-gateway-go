package usecase

import "context"

type GetRefundUseCase interface {
	Handle(ctx context.Context, input GetRefundInput) (GetRefundOutput, error)
}

type GetRefundInput struct{}

type GetRefundOutput struct{}
