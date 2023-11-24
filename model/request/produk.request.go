package request

import "github.com/google/uuid"

type CreateProduk struct {
	TokoID     uuid.UUID `db:"toko_id" json:"shop_id" validate:"required,uuid" gorm:"index;column:toko_id" form:"shop_id"`
	KategoriID uuid.UUID `db:"kategori_id" json:"category_id" validate:"required,uuid" gorm:"index;column:kategori_id" form:"category_id"`
	Nama       string    `json:"name" validate:"required" form:"name"`
	Stok       int       `json:"stock" validate:"required" form:"stock"`
	Harga      int       `json:"price" validate:"required" form:"price" `
	Image      string    `json:"photo" form:"photo"`
}

type UpdateProduk struct {
	TokoID     uuid.UUID `db:"toko_id" json:"shop_id" gorm:"index;column:toko_id" form:"shop_id"`
	KategoriID uuid.UUID `db:"kategori_id" json:"category_id" gorm:"index;column:kategori_id" form:"category_id"`
	Nama       string    `json:"name"  form:"name"`
	Stok       int       `json:"stock"  form:"stock"`
	Harga      int       `json:"price"  form:"price" `
	Image      string    `json:"photo" form:"photo"`
}
