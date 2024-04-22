package usecase

import "context"

type CreateRefundUseCase interface {
	Handle(ctx context.Context, input CreateRefundInput) (CreateRefundOutput, error)
}

type CreateRefundInput struct{}

type CreateRefundOutput struct{}
