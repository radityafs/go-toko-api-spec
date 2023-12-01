package routes

import (
	"go-toko/handler"

	"go-toko/middleware"

	"github.com/gofiber/fiber/v2"
)

func Init(r *fiber.App) {
	api := r.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/auth/login", handler.Login)
	v1.Post("/auth/refresh-token", handler.RefreshToken)

	v1.Get("/categories", middleware.Auth, handler.GetProductCategoriesShop)
	v1.Get("/Products", middleware.Auth, handler.GetProductsShop)
	v1.Get("/Top/Transactions", middleware.Auth, handler.GetTopTransactionByShop)
	v1.Get("/Transactions", middleware.Auth, handler.GetTransactionsByShop)
	v1.Get("/Transactions/:id", middleware.Auth, handler.GetDetailTransactionByShop)
	v1.Post("/Transactions", middleware.Auth, handler.PostCreateTransaction)
}
