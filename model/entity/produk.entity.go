package entity

import "github.com/google/uuid"

type Produk struct {
	ID         uuid.UUID `db:"id" json:"id" gorm:"primaryKey"`
	Toko       Toko      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TokoID     uuid.UUID `db:"toko_id" json:"shop_id" validate:"required,uuid" gorm:"index;column:toko_id" form:"shop_id"`
	Kategori   Kategori  `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	KategoriID uuid.UUID `db:"kategori_id" json:"category_id" validate:"required,uuid" gorm:"index;column:kategori_id" form:"category_id"`
	Nama       string    `json:"name" validate:"required" form:"name"`
	Harga      int       `json:"price" validate:"required" form:"price"`
	Stok       int       `json:"stock" validate:"required" form:"stock"`
	Image      string    `json:"photo" form:"photo" validate:"required"`
	Terjual    int       `json:"sold" form:"sold"`
	Time
}

func (Produk) TableName() string {
	return "produk"
}
