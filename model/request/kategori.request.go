package request

import "github.com/google/uuid"

type CreateKategori struct {
	TokoID uuid.UUID `db:"toko_id" json:"toko_id" validate:"required,uuid" gorm:"index;column:toko_id" form:"toko_id"`
	Nama   string    `json:"nama" validate:"required" form:"nama"`
}
