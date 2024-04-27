package gorm

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt int64
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Model) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	base.ID = uuid
	return
}
