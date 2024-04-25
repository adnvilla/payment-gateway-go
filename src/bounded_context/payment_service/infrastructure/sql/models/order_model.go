package models

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/gorm"
	uuid "github.com/satori/go.uuid"
)

type CreateOrder struct {
	gorm.Model
	Amount       string
	Currency     string
	ProviderType int

	CreateOrderProvider CreateOrderProvider `gorm:"foreignKey:CreateOrderID"`
}

type CreateOrderProvider struct {
	gorm.Model
	CreateOrderID uuid.UUID `gorm:"type:uuid"`
	ProviderType  int
	Payload       string
}
