package models

import (
	"github.com/adnvilla/payment-gateway-go/src/pkg/gorm"
	uuid "github.com/satori/go.uuid"
)

type CaptureOrder struct {
	gorm.Model
	Amount       string
	Currency     string
	ProviderType int

	CaptureOrderProvider CaptureOrderProvider `gorm:"foreignKey:CaptureOrderID"`
}

type CaptureOrderProvider struct {
	gorm.Model
	CaptureOrderID    uuid.UUID `gorm:"type:uuid"`
	ProviderCaptureID string
	ProviderType      int
	Payload           string
}
