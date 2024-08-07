package models

import (
	"gorm.io/gorm"
	"time"
)

// Base model. Not involved in migration
type Base struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP AT TIME ZONE 'UTC'" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone;default:CURRENT_TIMESTAMP AT TIME ZONE 'UTC'" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp without time zone" json:"-"`
}
