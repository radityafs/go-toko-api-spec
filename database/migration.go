package database

import (
	"fmt"
	"go-toko/model/entity"
)

func Migrate() {
	err := DB.AutoMigrate(
		entity.Toko{},
		entity.User{},
		entity.Kategori{},
		entity.Produk{},
		entity.Penjualan{},
		entity.DetailPenjualan{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated")
}
