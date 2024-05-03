package usecase

import (
	"context"
	"testing"

	mockRepository "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/repository/mock"
	mockService "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCaptureOrderUseCase(t *testing.T) {
	ctx := context.Background()
	mockService := mockService.NewMockOrderProviderService(t)
	mockRepository := mockRepository.NewMockOrderRepository(t)

	id := uuid.NewV4()
	input := CaptureOrderInput{}

	i := vo.CreateOrderDetail{
		OrderId: "orderproviderid",
	}

	info := vo.CaptureOrder{
		OrderId:      i.OrderId,
		ProviderType: input.ProviderType,
	}
	infoDetail := vo.CaptureOrderDetail{}

	mockRepository.EXPECT().CaptureOrder(ctx, infoDetail).Return(id, nil).Once()
	mockRepository.EXPECT().GetOrderProvider(ctx, input.OrderId).Return(i, nil).Once()
	mockService.EXPECT().CaptureOrder(ctx, info).Return(infoDetail, nil)

	u := NewCaptureOrderUseCase(mockService, mockRepository)

	out, err := u.Handle(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestGetOrderUseCase(t *testing.T) {
	ctx := context.Background()
	mockRepository := mockRepository.NewMockOrderRepository(t)

	id := uuid.NewV4()
	input := GetOrderInput{
		Id: id,
	}

	info := vo.CreateOrder{
		Id: id,
	}

	mockRepository.EXPECT().GetOrder(ctx, id).Return(info, nil).Once()

	u := NewGetOrderUseCase(mockRepository)

	out, err := u.Handle(ctx, input)

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestCreateOrderUseCase(t *testing.T) {
	ctx := context.Background()
	mockService := mockService.NewMockOrderProviderService(t)
	mockRepository := mockRepository.NewMockOrderRepository(t)

	input := CreateOrderInput{}

	id := uuid.NewV4()
	info := vo.CreateOrder{
		ProviderType: input.ProviderType,
	}
	infoDetail := vo.CreateOrderDetail{}

	mockRepository.EXPECT().CreateOrder(ctx, infoDetail).Return(id, nil).Once()
	mockService.EXPECT().CreateOrder(ctx, info).Return(infoDetail, nil)

	u := NewCreateOrderUseCase(mockService, mockRepository)

	out, err := u.Handle(ctx, CreateOrderInput{})

	assert.NoError(t, err)
	assert.NotNil(t, out)
}

func TestCreateRefundUseCase(t *testing.T) {
	ctx := context.Background()
	mockService := mockService.NewMockOrderProviderService(t)
	mockRepository := mockRepository.NewMockOrderRepository(t)

	id := uuid.NewV4()
	input := CreateRefundInput{
		CaptureOrderId: id,
	}

	i := vo.CaptureOrderDetail{
		CaptureOrderId: "orderproviderid",
	}

	info := vo.CreateRefundOrder{
		CaptureOrderId: i.CaptureOrderId,
		ProviderType:   input.ProviderType,
	}
	infoDetail := vo.CreateRefundDetail{}

	mockRepository.EXPECT().CreateRefund(ctx, infoDetail).Return(id, nil).Once()
	mockRepository.EXPECT().GetCaptureOrderProvider(ctx, input.CaptureOrderId).Return(i, nil).Once()
	mockService.EXPECT().CreateRefund(ctx, info).Return(infoDetail, nil)

	u := NewCreateRefundUseCase(mockService, mockRepository)

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
