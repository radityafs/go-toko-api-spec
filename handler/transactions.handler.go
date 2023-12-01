package handler

import (
	"errors"
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

		response.Data = nil
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

func PostCreateTransaction(ctx *fiber.Ctx) error {
	var shop_id = ctx.Locals("shop_id")
	var cashier_id = ctx.Locals("user_id")
	var request request.Sales
	var Transaction response.CreateTransaction


	// float to uint64
	shopId := uint64(shop_id.(float64))
	cashierId := uint64(cashier_id.(float64))

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"data":    nil,
		})
	}

	trx := database.DB.Transaction(func(tx *gorm.DB) error {
		tx.Begin()

		var sales entity.Sales
		var orderId = uuid.New().String()

		sales.OrderID = orderId
		sales.ShopID = shopId
		sales.CashierID = cashierId
		sales.PaymentType = request.PaymentType

		err := tx.Create(&sales).Error;
		if(err != nil) {
			tx.Rollback()
			return err
		}

		Transaction.Data.Id = sales.ID
		Transaction.Data.Invoice = orderId

		for _, product := range request.Product{
			var productData entity.Product
			var salesDetail entity.SalesDetail

			database.DB.Preload("ProductsCategory").Find(&productData, "id = ?", product.ID)

			if(productData.ID == 0) {
				tx.Rollback()
				return errors.New("product not found")
			}

			if(productData.Quantity < product.Quantity) {
				tx.Rollback()
				return errors.New("product quantity not enough")
			}

			Transaction.Data.Product = append(Transaction.Data.Product, response.DetailProduct{
				Id: productData.ID,
				Name: productData.Name,
				Price: productData.PriceSell,
				Quantity: product.Quantity,
				Total: product.Quantity * productData.PriceSell,
			})

			Transaction.Data.TotalPayment += product.Quantity * productData.PriceSell

			// update product quantity
			productData.Quantity = productData.Quantity - product.Quantity
			err := tx.Save(&productData).Error
			
			if(err != nil) {
				tx.Rollback()
				return err
			}

			// create sales detail
			salesDetail.SalesID = sales.ID
			salesDetail.ProductID = productData.ID
			salesDetail.ShopID = shopId
			salesDetail.Name = productData.Name
			salesDetail.Quantity = product.Quantity
			salesDetail.Category = productData.ProductsCategory.Name
			salesDetail.Price = productData.PriceSell
			salesDetail.Total = product.Quantity * productData.PriceSell

			errs := tx.Create(&salesDetail).Error

			if(errs != nil) {
				tx.Rollback()
				return errs
			}
		}

		if(sales.PaymentType == entity.Cash) {
			sales.Status = entity.Paid
		}else{
			sales.Status = entity.Unpaid
			Transaction.Data.PaymentRef = uuid.New().String()
		}

		if(sales.TotalPaid < sales.TotalBill) {
			tx.Rollback()
			return errors.New("total payment less than total bill")
		}

		Transaction.Data.Change = request.TotalPayment - sales.TotalBill
		Transaction.Data.PaymentMethod = string(sales.PaymentType)
		Transaction.Data.PaymentStatus = string(sales.Status)
		Transaction.Data.Total = Transaction.Data.TotalPayment

		sales.TotalPaid = request.TotalPayment
		sales.TotalItem = len(request.Product)
		sales.TotalBill = Transaction.Data.TotalPayment

		errSave := tx.Save(&sales).Error

		if(errSave != nil) {
			tx.Rollback()
			return errSave
		}

		return nil
	})

	if(trx != nil) {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": trx.Error(),
			"data":    nil,
		})
	}

	Transaction.Status = 200
	Transaction.Success = true
	Transaction.Message = "Berhasil membuat transaksi"
	return ctx.Status(Transaction.Status).JSON(Transaction)
}