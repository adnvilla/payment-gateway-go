package usecase

import (
	"context"
	"testing"

	mockRepository "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository/mock"
	mockService "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/stretchr/testify/assert"
)

func TestCaptureOrderUseCase(t *testing.T) {
	ctx := context.Background()
	mockService := mockService.NewMockOrderProviderService(t)
	mockRepository := mockRepository.NewMockOrderRepository(t)

	input := CaptureOrderInput{}

	info := vo.CaptureOrder{
		OrderId:      input.OrderId,
		ProviderType: input.ProviderType,
	}
	infoDetail := vo.CaptureOrderDetail{}

	mockRepository.EXPECT().CaptureOrder(ctx, infoDetail).Return(nil).Once()
	mockService.EXPECT().CaptureOrder(ctx, info).Return(infoDetail, nil)

	u := NewCaptureOrderUseCase(mockService, mockRepository)

	out, err := u.Handle(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestCreateOrderUseCase(t *testing.T) {
	ctx := context.Background()
	mockService := mockService.NewMockOrderProviderService(t)
	mockRepository := mockRepository.NewMockOrderRepository(t)

	input := CreateOrderInput{}

	info := vo.CreateOrder{
		ProviderType: input.ProviderType,
	}
	infoDetail := vo.CreateOrderDetail{}

	mockRepository.EXPECT().CreateOrder(ctx, infoDetail).Return(nil).Once()
	mockService.EXPECT().CreateOrder(ctx, info).Return(infoDetail, nil)

	u := NewCreateOrderUseCase(mockService, mockRepository)

	out, err := u.Handle(ctx, CreateOrderInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestCreateRefundUseCase(t *testing.T) {
	ctx := context.Background()

	input := CreateRefundInput{}

	u := NewCreateRefundUseCase()
	out, err := u.Handle(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestGetRefundUseCase(t *testing.T) {
	ctx := context.Background()

	input := GetRefundInput{}

	u := NewGetRefundUseCase()
	out, err := u.Handle(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, out)
}
