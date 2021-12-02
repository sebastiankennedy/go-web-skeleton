package models

import (
	"github.com/sebastiankennedy/go-web-skeleton/pkg/types"
	"time"
)

type Model struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL"`
}

func (b Model) GetStringID() string {
	return types.Uint64ToString(b.ID)
}