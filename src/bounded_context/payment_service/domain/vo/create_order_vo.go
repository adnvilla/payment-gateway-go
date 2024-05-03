package vo

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
)

type CreateOrder struct {
	Id           uuid.UUID
	Amount       string
	Currency     string
	CreatedAt    int64
	ProviderType shared_domain.ProviderType
}

type CreateOrderDetail struct {
	Id           uuid.UUID
	OrderId      string
	ProviderType shared_domain.ProviderType
	Amount       string
	Currency     string
	CreatedAt    int64
	Payload      string
}
