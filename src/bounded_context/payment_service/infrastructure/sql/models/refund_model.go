package models

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/gorm"
	uuid "github.com/satori/go.uuid"
)

type Refund struct {
	gorm.Model
	Amount       string
	Currency     string
	ProviderType int

	RefundProvider RefundProvider `gorm:"foreignKey:RefundId"`
}

type RefundProvider struct {
	gorm.Model
	RefundId         uuid.UUID `gorm:"type:uuid"`
	ProviderRefundID string
	ProviderType     int
	Payload          string
}
