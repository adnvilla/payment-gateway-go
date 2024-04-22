package usecase

import "context"

type CreateOrderUseCase interface {
	Handle(ctx context.Context, input CreateOrderInput) (CreateOrderOutput, error)
}

type CreateOrderInput struct{}

type CreateOrderOutput struct{}
