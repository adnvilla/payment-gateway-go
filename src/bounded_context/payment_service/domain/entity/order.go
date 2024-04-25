package entity

import "github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"

type Order struct {
	Id           string
	Amount       float32
	Currency     string
	CreatedAt    int64
	ProviderType shared_domain.ProviderType
}

type OrderProvider struct {
	Id           string
	OrderId      string
	Payload      string
	ProviderType shared_domain.ProviderType
	CreatedAt    int64
}
