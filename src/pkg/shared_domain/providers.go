package shared_domain

type ProviderType int

const (
	ProviderType_Stripe ProviderType = iota + 1
	ProviderType_Paypal
)
