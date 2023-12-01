package response

import "go-toko/model/entity"

type Products struct {
	Status  int         `json:"-"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    []ProductsData `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type ProductsData struct {
	Id int `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Stock int `gorm:"column:quantity" json:"stock"`
	Price int `gorm:"column:price_sell" json:"price"`
	Photo string `gorm:"column:images" json:"photo"`
	ProductsCategory *entity.ProductsCategory `gorm:"foreignKey:CategoryID" json:"category"`
	CategoryID int `gorm:"column:category_id" json:"-"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}
