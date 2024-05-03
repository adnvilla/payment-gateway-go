package vo

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	uuid "github.com/satori/go.uuid"
)

type CreateRefundOrder struct {
	Id             uuid.UUID
	CaptureOrderId string
	ProviderType   shared_domain.ProviderType
}

type CreateRefundDetail struct {
	Id            uuid.UUID
	RefundOrderId string
	ProviderType  shared_domain.ProviderType
	Payload       string
}
