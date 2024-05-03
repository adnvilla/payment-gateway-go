package vo

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
)

type CaptureOrder struct {
	Id           uuid.UUID
	OrderId      string
	ProviderType shared_domain.ProviderType
}

type CaptureOrderDetail struct {
	Id             uuid.UUID
	CaptureOrderId string
	ProviderType   shared_domain.ProviderType
	Payload        string
}
