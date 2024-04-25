package vo

import "github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"

type CreateOrder struct {
	Amount       string
	Currency     string
	ProviderType shared_domain.ProviderType
}

type CreateOrderDetail struct {
	Id           string
	OrderId      string
	ProviderType shared_domain.ProviderType
	Amount       string
	Currency     string
	CreatedAt    int64
	Payload      string
}

type CaptureOrder struct {
	OrderId      string
	ProviderType shared_domain.ProviderType
}

type CaptureOrderDetail struct {
	Id           string
	ProviderType shared_domain.ProviderType
	Amount       string
	Currency     string
	CreatedAt    int64
	Payload      string
}
