package entity

import (
	"time"

	"gorm.io/gorm"
)

type Time struct {
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index; column:deleted_at"`
}
