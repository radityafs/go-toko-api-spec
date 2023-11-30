package routes

import (
	"go-toko/config"
	"go-toko/handler"

	"go-toko/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r *fiber.App) {
	r.Static("/media", config.ProjectRootPath+"/public")
	api := r.Group("/api")
	v1 := api.Group("/v1")

	// cashier := v1.Group("/cashier", middleware.Auth)

	//! *************************
	//! * Auth Routes 					*
	//! *************************
	v1.Post("/auth/login", handler.Login)
	v1.Post("/auth/refresh-token", handler.RefreshToken)

	//! *************************
	//! * User Routes 					*
	//! *************************
	// v1.Get("/users", middleware.Auth, middleware.IsAdmin, handler.GetUser)
	// v1.Post("/users", middleware.Auth, middleware.IsAdmin, handler.CreateUser)

	//! *************************
	//! * Toko Routes 					*
	//! *************************
	// v1.Get("/shops", middleware.Auth, middleware.IsAdmin, handler.GetToko)
	// v1.Get("/shops/:id", middleware.Auth, middleware.IsAdmin, handler.GetOneToko)
	// v1.Post("/shops", middleware.Auth, middleware.IsAdmin, handler.CreateToko)
	// v1.Put("/shops/:id", middleware.Auth, middleware.IsAdmin, handler.UpdateToko)
	// v1.Delete("/shops/:id", middleware.Auth, middleware.IsAdmin, handler.DeleteToko)

	//! *************************
	//! * Produk Routes 				*
	//! *************************
	// v1.Get("/products", middleware.Auth, handler.GetProduk)
	// v1.Put("/products/:id", middleware.Auth, middleware.IsAdmin, handler.UpdateProduk)
	// v1.Post("/products", middleware.Auth, middleware.IsAdmin, handler.CreateProduk)

	//! ***************************
	//! * Kategori Routes 				*
	//! ***************************
	v1.Get("/categories", middleware.Auth, handler.GetProductCategoriesShop)
	v1.Get("/Products", middleware.Auth, handler.GetProductsShop)
	v1.Get("/Top/Transactions", middleware.Auth, handler.GetTopTransactionByShop)
	v1.Get("/Transactions", middleware.Auth, handler.GetTransactionsByShop)
	v1.Get("/Transactions/:id", middleware.Auth, handler.GetDetailTransactionByShop)
	// admin.Get("/kategori", handler.GetKategori)
	// admin.Post("/kategori", handler.CreateKategori)

	// owner.Get("/kategori", handler.GetKategori)
	// owner.Post("/kategori", handler.CreateKategori)

	//! ***************************
	//! * Penjualan Routes 				*
	//! ***************************
	// v1.Get("/transactions", middleware.Auth, handler.GetPenjualan)
	// v1.Post("/transactions", middleware.Auth, handler.CreatePenjualan)
	// v1.Get("/transactions/top", middleware.Auth, handler.GetTopData)
	// v1.Get("/transactions/:id", middleware.Auth, handler.GetOnePenjualan)
}
