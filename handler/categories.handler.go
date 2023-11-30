package handler

import (
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/response"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductCategoriesShop(ctx *fiber.Ctx) error {
	page:= ctx.Query("page", "1")
	limit := ctx.Query("limit", "10")

	pageInt, errPageInt := strconv.Atoi(page)
	limitInt, errLimitInt := strconv.Atoi(limit)

	if errPageInt != nil || errLimitInt != nil || pageInt < 1 || limitInt < 1 {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Query page dan limit harus berupa angka dan lebih dari 0",
			"data":    nil,
		})
	}

	// Variable
	var response response.Categories

	// Query
	tx := database.DB.Model(&entity.ProductsCategory{}).Where("shop_id = ?", ctx.Locals("shop_id"))

	tx.Count(&response.Pagination.TotalData)
	tx.Limit(limitInt).Offset((pageInt-1)*limitInt).Scan(&response.Data)

	// Pagination
	response.Pagination.CurrentPage = int64(pageInt)
	response.Pagination.TotalPage = int64(math.Ceil(float64(response.Pagination.TotalData) / float64(limitInt)))
	response.Pagination.PerPage = int64(limitInt)

	if(response.Pagination.TotalData == 0) {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
	}else{
		response.Status = 200
		response.Message = "Berhasil mengambil data kategori"
	}

	return ctx.Status(response.Status).JSON(response)
}