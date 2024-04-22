package usecase

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
)

type CaptureOrderUseCase interface {
	Handle(ctx context.Context, input CaptureOrderInput) (CaptureOrderOutput, error)
}

type CaptureOrderInput struct{}

type CaptureOrderOutput struct{}

type captureOrderUseCase struct{}

func NewCaptureOrderUseCase() use_case.UseCase[CaptureOrderInput, CaptureOrderOutput] {
	u := new(captureOrderUseCase)
	return u
}

func (u *captureOrderUseCase) Handle(ctx context.Context, input CaptureOrderInput) (CaptureOrderOutput, error) {
	output := CaptureOrderOutput{}
	return output, nil
}
