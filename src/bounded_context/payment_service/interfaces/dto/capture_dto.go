package dto

type CaptureOrderRequest struct {
	OrderId      string `json:"order_id"`
	ProviderType int    `json:"provider_type"`
}

type CaptureOrderResponse struct {
	Id           string `json:"id"`
	ProviderType int    `json:"provider_type"`
	Amount       string `json:"amount"`
	Currency     string `json:"currency"`
	CreatedAt    int    `json:"created_at"`
}
