package handler

import (
	"fmt"
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/response"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductsShop(ctx *fiber.Ctx) error {
	page:= ctx.Query("page", "1")
	limit := ctx.Query("limit", "10")
	id_category := ctx.Query("id_category", "")
	search := ctx.Query("search", "")
	sortBy := ctx.Query("sort_by", "")
	sort := ctx.Query("sort", "")

	pageInt, errPageInt := strconv.Atoi(page)
	limitInt, errLimitInt := strconv.Atoi(limit)

	var response = response.Products{}

	if errPageInt != nil || errLimitInt != nil || pageInt < 1 || limitInt < 1 {
		response.Status = 400
		response.Message = "Query page dan limit harus berupa angka dan lebih dari 0"
		response.Success = false

		return ctx.Status(response.Status).JSON(response)
	}

	if(sort != "" && sort != "asc" && sort != "desc") {
		response.Status = 400
		response.Message = "Query sort harus berupa asc atau desc"
		response.Success = false

		return ctx.Status(response.Status).JSON(response)
	}

	if(sortBy != "" && sortBy != "id" && sortBy != "created_at" && sortBy != "name") {
		response.Status = 400
		response.Message = "Query sort_by harus berupa id, created_at, atau name"
		response.Success = false

		return ctx.Status(response.Status).JSON(response)
	}


	tx := database.DB.Preload("ProductsCategory").
	Model(&entity.Product{}).
	Where("shop_id = ?", ctx.Locals("shop_id")).
	Where("quantity > ?", 0)
	
	tx.Count(&response.Pagination.TotalData)

	if(search != "") { 
		tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", search))
		tx.Count(&response.Pagination.TotalData)
	}

	if(id_category != "") { 
		tx.Where("category_id = ?", id_category)
		tx.Count(&response.Pagination.TotalData)
	}

	if(sortBy != "" && sort != "") { tx.Order(fmt.Sprintf("%s %s", sortBy, sort)) }
	tx.Limit(limitInt).Offset((pageInt-1)*limitInt).Find(&response.Data)

	response.Pagination.TotalPage = int64(math.Ceil(float64(response.Pagination.TotalData) / float64(limitInt)))
	response.Pagination.PerPage = int64(limitInt)
	response.Pagination.CurrentPage = int64(pageInt)

	if(response.Pagination.TotalData == 0) {
		response.Status = 404
		response.Message = "Data tidak ditemukan"
		response.Success = false
	}else{
		response.Status = 200
		response.Message = "Berhasil mengambil data produk"
		response.Success = true
	}

	return ctx.Status(response.Status).JSON(response)
}