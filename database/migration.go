package database

import (
	"fmt"
	"go-toko/model/entity"
)

func Migrate() {

	err := DB.AutoMigrate(
		&entity.User{},
		&entity.UserRole{},
		&entity.Shops{},
		&entity.ProductsBrand{},
		&entity.ProductsCategory{},
		&entity.Product{},
		&entity.Sales{},
		&entity.SalesDetail{},
		&entity.SalesPayment{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated")
}
