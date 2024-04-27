package dto

type GetOrderRequest struct {
	Id string `json:"id"`
}

type GetOrderResponse struct {
	Id           string `json:"id"`
	ProviderType int    `json:"provider_type"`
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	CreatedAt    int    `json:"created_at"`
}
