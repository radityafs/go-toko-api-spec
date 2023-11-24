package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid" gorm:"primaryKey"`
	Email    string    `json:"email" validate:"required,email" gorm:"unique"`
	Password string    `json:"-" gorm:"column:password"`
	Nama     string    `json:"name" validate:"required"`
	Toko     Toko      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TokoID   uuid.UUID `json:"shop_id" validate:"required,uuid" gorm:"column:toko_id;index"`
	Role     string    `json:"role" validate:"oneof=admin cashier owner"`
	Time
}

func (User) TableName() string {
	return "user"
}
