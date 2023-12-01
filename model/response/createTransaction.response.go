package response

type CreateTransaction struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data   CreateTransactionData `json:"data"`
}

type CreateTransactionData struct {
	Id uint64 `gorm:"column:id" json:"id"`
	Invoice string `gorm:"column:order_id" json:"invoice"`
	PaymentMethod string `gorm:"column:payment_type" json:"payment_method"`
	PaymentStatus string `gorm:"column:status" json:"status"`
	PaymentRef string `json:"payment_ref"`
	Total int `gorm:"column:total_bill" json:"total"`
	TotalPayment int `gorm:"column:total_paid" json:"total_payment"`
	Change int `gorm:"column:change" json:"change"`
	Product []DetailProduct `json:"product"`
}