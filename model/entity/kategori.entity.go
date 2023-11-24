package entity

import "github.com/google/uuid"

type Kategori struct {
	ID      uuid.UUID `db:"id" json:"id" gorm:"primaryKey"`
	Nama    string    `db:"nama" json:"name" validate:"required"`
	Toko    Toko      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TokoID  uuid.UUID `db:"toko_id" json:"toko_id" validate:"required,uuid" gorm:"index;column:toko_id" form:"toko_id"`
	Produk  []Produk  `json:"-" gorm:"foreignKey:KategoriID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Terjual int       `json:"sold" form:"sold"`
	Time
}

func (Kategori) TableName() string {
	return "kategori"
}
