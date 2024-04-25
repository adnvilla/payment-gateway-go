package dto

type CreateOrderRequest struct {
	ProviderType int    `json:"provider_type"`
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
}

type CreateOrderResponse struct{}
