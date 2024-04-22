package usecase

import "context"

type CaptureOrderUseCase interface {
	Handle(ctx context.Context, input CaptureOrderInput) (CaptureOrderOutput, error)
}

type CaptureOrderInput struct{}

type CaptureOrderOutput struct{}
