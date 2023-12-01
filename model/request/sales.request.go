package request

import "go-toko/model/entity"

type Sales struct {
	Product []ProductData `json:"product"`
	TotalPayment int `json:"total_payment"`
	PaymentType entity.PaymentType `json:"payment_type"`
}

type ProductData struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
}