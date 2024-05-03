package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	serviceMock "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service/mock"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/vo"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetService(t *testing.T) {
	s := service.NewCreateOrderService(nil)
	assert.NotNil(t, s)
}

func TestGetServiceCreateOrder(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType_Stripe).Return(mockProvider, nil).Once()
	mockProvider.EXPECT().CreateOrder(ctx, mock.Anything).Return(vo.CreateOrderDetail{}, nil).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CreateOrder(ctx, vo.CreateOrder{
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestGetServiceCreateOrderUnsupportedProvider(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType(100)).Return(mockProvider, errors.New("unsupported provider")).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CreateOrder(ctx, vo.CreateOrder{
		ProviderType: shared_domain.ProviderType(100),
	})

	assert.Error(t, err)
	assert.NotNil(t, r)
}

func TestGetServiceCaptureOrder(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType_Stripe).Return(mockProvider, nil).Once()
	mockProvider.EXPECT().CaptureOrder(ctx, mock.Anything).Return(vo.CaptureOrderDetail{}, nil).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CaptureOrder(ctx, vo.CaptureOrder{
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestGetServiceCaptureOrderUnsupportedProvider(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType(100)).Return(mockProvider, errors.New("unsupported provider")).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CaptureOrder(ctx, vo.CaptureOrder{
		ProviderType: shared_domain.ProviderType(100),
	})

	assert.Error(t, err)
	assert.NotNil(t, r)
}

func TestGetServiceCreateRefund(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType_Stripe).Return(mockProvider, nil).Once()
	mockProvider.EXPECT().CreateRefund(ctx, mock.Anything).Return(vo.CreateRefundDetail{}, nil).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CreateRefund(ctx, vo.CreateRefundOrder{
		ProviderType: shared_domain.ProviderType_Stripe,
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)
}

func TestGetServiceCreateRefundUnsupportedProvider(t *testing.T) {
	ctx := context.Background()
	mockFactory := serviceMock.NewMockGetProviderService(t)
	mockProvider := serviceMock.NewMockOrderProviderService(t)

	mockFactory.EXPECT().GetProviderClient(ctx, shared_domain.ProviderType(100)).Return(mockProvider, errors.New("unsupported provider")).Once()
	s := service.NewCreateOrderService(mockFactory)

	r, err := s.CreateRefund(ctx, vo.CreateRefundOrder{
		ProviderType: shared_domain.ProviderType(100),
	})

	assert.Error(t, err)
	assert.NotNil(t, r)
}
