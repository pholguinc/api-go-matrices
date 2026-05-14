package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MatrixRecord struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	UserID    string         `gorm:"type:varchar(36);index" json:"user_id"`
	Input     [][]float64    `gorm:"type:jsonb" json:"input"`
	Q         [][]float64    `gorm:"type:jsonb" json:"q"`
	R         [][]float64    `gorm:"type:jsonb" json:"r"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *MatrixRecord) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}
