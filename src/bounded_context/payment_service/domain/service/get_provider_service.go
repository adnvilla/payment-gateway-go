package service

import (
	"context"

	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
)

type GetProviderService interface {
	GetProviderClient(ctx context.Context, provider shared_domain.ProviderType) (OrderProviderService, error)
}
