package response

import "go-toko/model/entity"

type Transaction struct {
	Status int `json:"-"`
	Success bool `json:"success"`
	Message string `json:"message"`
	Data DetailTransaction `json:"data"`
}

type DetailTransaction struct {
	Id uint64 `gorm:"column:id" json:"id"`
	Invoice string `gorm:"column:order_id" json:"invoice"`
	PaymentMethod string `gorm:"column:payment_type" json:"payment_method"`
	PaymentStatus entity.PaymentStatus `gorm:"column:status" json:"status"`
	Total int `gorm:"column:total_bill" json:"total"`
	TotalPayment int `gorm:"column:total_paid" json:"total_payment"`
	Change int `gorm:"column:change" json:"change"`
	Product []DetailProduct `json:"product"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

type DetailProduct struct {
	Id uint64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Price int `gorm:"column:price" json:"price"`
	Quantity int `gorm:"column:quantity" json:"quantity"`
	Total int `gorm:"column:total" json:"total"`
}