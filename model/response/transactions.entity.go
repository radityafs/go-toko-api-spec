package response

import "go-toko/model/entity"

type Transactions struct {
	Status  int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []DetailTransactions `json:"data"`
	Summary SummaryTransactions `json:"summary"`
	Pagination Pagination `json:"pagination"`
}

type DetailTransactions struct {
	Id int `gorm:"column:id" json:"id"`
	Invoice string `gorm:"column:order_id" json:"invoice"`
	PaymentMethod entity.PaymentType `gorm:"column:payment_type" json:"payment_method"`
	Status entity.PaymentStatus `gorm:"column:status" json:"status"`
	Total int `gorm:"column:total_bill" json:"total"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

type SummaryTransactions struct {
	TotalRevenue int `gorm:"column:total_revenue" json:"total_revenue"`
	TotalSales int `gorm:"column:total_sales" json:"total_sales"`
}