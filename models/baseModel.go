package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time      `gorm:"column:created_at;null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}
