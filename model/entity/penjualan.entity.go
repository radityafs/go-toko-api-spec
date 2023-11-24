package entity

import "github.com/google/uuid"

type Penjualan struct {
	ID              uuid.UUID         `db:"id" json:"id" gorm:"primaryKey"`
	User            User              `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID          uuid.UUID         `db:"user_id" json:"-" gorm:"index;column:user_id"`
	OrderID         string            `db:"order_id" json:"invoice"`
	QrisID          string            `db:"qris_id" json:"qris_id"`
	Status          string            `db:"status" json:"status"`
	Toko            Toko              `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TokoID          uuid.UUID         `db:"toko_id" json:"-" gorm:"index;column:toko_id"`
	Total           int               `db:"total" json:"total"`
	IipePembayaran  string            `db:"tipe_pembayaran" json:"payment_method"`
	TotalPembayaran int               `db:"total_pembayaran" json:"-"`
	JumlahItem      int               `db:"jumlah_item" json:"-"`
	Kembalian       int               `db:"kembalian" json:"-"`
	DetailPenjualan []DetailPenjualan `json:"-" gorm:"foreignKey:PenjualanID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Time
}

type DetailPenjualan struct {
	ID          uuid.UUID `db:"id" json:"transaction_detail_id" gorm:"primaryKey"`
	Quantity    int       `db:"quantity" json:"qty"`
	Produk      Produk    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProdukName  string    `db:"produk_name" json:"name"`
	ProdukID    uuid.UUID `db:"produk_id" json:"id" validate:"required,uuid" gorm:"index;column:produk_id"`
	Harga       int       `db:"harga" json:"price"`
	Subtotal    int       `db:"subtotal" json:"subtotal"`
	Penjualan   Penjualan `json:"-" gorm:"foreignKey:PenjualanID"`
	PenjualanID uuid.UUID `db:"penjualan_id" json:"-" gorm:"index;column:penjualan_id"`
}

type OnePenjualan struct {
	ID              uuid.UUID         `db:"id" json:"id" gorm:"primaryKey"`
	User            User              `json:"kasir" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID          uuid.UUID         `db:"user_id" json:"-" gorm:"index;column:user_id"`
	OrderID         string            `db:"order_id" json:"invoice"`
	QrisID          string            `db:"qris_id" json:"qris_id"`
	Status          string            `db:"status" json:"status"`
	Toko            Toko              `json:"toko" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TokoID          uuid.UUID         `db:"toko_id" json:"toko_id" gorm:"index;column:toko_id"`
	Total           int               `db:"total" json:"total"`
	IipePembayaran  string            `db:"tipe_pembayaran" json:"payment_method"`
	TotalPembayaran int               `db:"total_pembayaran" json:"total_payment"`
	JumlahItem      int               `db:"jumlah_item" json:"total_item"`
	Kembalian       int               `db:"kembalian" json:"change"`
	DetailPenjualan []DetailPenjualan `json:"products" gorm:"foreignKey:PenjualanID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Time
}

func (Penjualan) TableName() string {
	return "penjualan"
}

func (DetailPenjualan) TableName() string {
	return "detail_penjualan"
}
