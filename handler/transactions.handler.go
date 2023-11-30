package handler

import (
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func GetTopTransactionByShop(ctx *fiber.Ctx) error {
	page:= ctx.Query("page", "1")
	limit := ctx.Query("limit", "10")
	typeBy := ctx.Query("type", "category")

	start_date := ctx.Query("start_date", "2023-11-20")
	end_date := ctx.Query("end_date", "2023-11-29")

	pageInt, errPageInt := strconv.Atoi(page)
	limitInt, errLimitInt := strconv.Atoi(limit)

	var response response.Analytics


	if errPageInt != nil || errLimitInt != nil || pageInt < 1 || limitInt < 1 {	
		response.Status = 400
		response.Message = "Query page dan limit harus berupa angka dan lebih dari 0"
		response.Success = false
		return ctx.Status(response.Status).JSON(response)
	}
	

	if(typeBy != "category" && typeBy != "product") {
		response.Status = 400
		response.Message = "Query type harus berupa category atau product"
		response.Success = false
		return ctx.Status(response.Status).JSON(response)
	}
	

	if(start_date == "" && end_date == "") {
		response.Status = 400
		response.Message = "Query start_date dan end_date harus diisi"
		response.Success = false
		return ctx.Status(response.Status).JSON(response)
	}

	tx := database.DB.Model(&entity.SalesDetail{}).
	Where("shop_id = ?", ctx.Locals("shop_id")).
	Where("created_at BETWEEN ? AND ?", start_date, end_date)

	tx.Count(&response.Pagination.TotalData)

	tx.Group(typeBy).
	Select("SUM(quantity) as total_sales, SUM(total) as total_revenue, "+typeBy+" as name").
	Order("total_sales desc").
	Limit(limitInt).
	Offset((pageInt-1)*limitInt).
	Scan(&response)

	response.Pagination.CurrentPage = int64(pageInt)
	if(response.Pagination.TotalData == 0) {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
	}else{
		response.Pagination.TotalPage = int64(response.Pagination.TotalData/ int64(limitInt)) + 1
		response.Pagination.PerPage = int64(limitInt)
		response.Status = 200
		response.Message = "Berhasil mengambil data transaksi"
	}

	return ctx.Status(response.Status).JSON(response)
}

func GetTransactionsByShop(ctx *fiber.Ctx) error {
	page:= ctx.Query("page", "1")
	limit := ctx.Query("limit", "10")

	start_date := ctx.Query("start_date", "2023-11-20")
	end_date := ctx.Query("end_date", "2023-11-29")

	pageInt, errPageInt := strconv.Atoi(page)
	limitInt, errLimitInt := strconv.Atoi(limit)

	var response response.Transactions

	if errPageInt != nil || errLimitInt != nil || pageInt < 1 || limitInt < 1 {
		response.Status = 400
		response.Message = "Query page dan limit harus berupa angka dan lebih dari 0"
		response.Success = false

		return ctx.Status(response.Status).JSON(response)
	}

	if(start_date == "" && end_date == "") {
		response.Status = 400
		response.Message = "Query start_date dan end_date harus diisi"
		response.Success = false
		return ctx.Status(response.Status).JSON(response)
	}


	tx := database.DB.Model(&entity.Sales{}).
	Where("created_at BETWEEN ? AND ?", start_date, end_date).
	Where("shop_id = ?", ctx.Locals("shop_id"))
	
	tx.Count(&response.Pagination.TotalData)
	tx.Select("COUNT(*) as total_sales, SUM(total_bill) as total_revenue")
	tx.Limit(limitInt).Offset((pageInt-1)*limitInt).Scan(&response.Data)

	if(response.Pagination.TotalData == 0) {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
		response.Success = false
	}else{
		response.Pagination.CurrentPage = int64(pageInt)
		response.Pagination.PerPage = int64(limitInt)
		response.Pagination.TotalPage = int64(response.Pagination.TotalData/ int64(limitInt)) + 1

		response.Status = 200
		response.Message = "Berhasil mengambil data transaksi"
		response.Success = true
	}

	return ctx.Status(response.Status).JSON(response)
}

func GetDetailTransactionByShop(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var response response.Transaction

	database.DB.Model(&entity.Sales{}).Where("order_id = ?", id).Where("shop_id = ?", ctx.Locals("shop_id")).Scan(&response.Data)
	database.DB.Model(&entity.SalesDetail{}).Where("sales_id = ?", response.Data.Id).Scan(&response.Data.Product)

	if(response.Data.Id == 0) {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
		response.Success = false
	}else{
		response.Status = 200
		response.Message = "Berhasil mengambil data transaksi"
		response.Success = true
	}

	return ctx.Status(response.Status).JSON(response)
}
