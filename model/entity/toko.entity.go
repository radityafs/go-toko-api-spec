package entity

import "github.com/google/uuid"

type Toko struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" gorm:"primaryKey"`
	Nama     string     `json:"name" validate:"required"`
	Alamat   string     `json:"address"`
	Telepon  string     `json:"phone"`
	User     []User     `json:"-" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Produk   []Produk   `json:"-" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Kategori []Kategori `json:"-" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Time
}

type OneToko struct {
	ID       uuid.UUID  `json:"id" validate:"required,uuid" gorm:"primaryKey"`
	Nama     string     `json:"name" validate:"required"`
	Alamat   string     `json:"address"`
	Telepon  string     `json:"phone"`
	User     []User     `json:"user" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Produk   []Produk   `json:"products" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Kategori []Kategori `json:"categories" gorm:"foreignKey:TokoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Time
}

func (Toko) TableName() string {
	return "toko"
}
