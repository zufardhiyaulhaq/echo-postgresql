package types

import (
	"time"

	"gorm.io/gorm"
)

type Echo struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ID        string         `gorm:"type:varchar(255);not null" json:"id"`
	Echo      string         `gorm:"type:varchar(255);not null" json:"echo"`
}
