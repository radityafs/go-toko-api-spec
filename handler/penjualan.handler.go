package handler

import (
	"fmt"
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/utils"
	"math"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func parseDate(date string) string {
	// input: DD-MM-YYYY
	// output: YYYY-MM-DD

	split := strings.Split(date, "-")

	return fmt.Sprintf("%s-%s-%s", split[2], split[1], split[0])
}

func generateQrisID() string {
	// format: qris-<8 random number>
	// example: qris-12345678

	var qrisID string

	// generate latest qris id
	var penjualan []entity.Penjualan

	database.DB.
		Where("qris_id LIKE ?", fmt.Sprintf("qris-%%")).
		Find(&penjualan)

	if len(penjualan) == 0 {
		qrisID = fmt.Sprintf("qris-%08d", 1)
	} else {
		qrisID = fmt.Sprintf("qris-%08d", len(penjualan)+1)
	}

	return qrisID
}

func generateOrderID() string {
	// format: INV-<YYYYMMDD>-<number pad start 4>
	// example: INV-20210901-0001

	var orderID string

	// generate date
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	dateStr := fmt.Sprintf("%d%02d%02d", year, month, day)

	// generate latest order id
	var penjualan []entity.Penjualan
	database.DB.
		Where("order_id LIKE ?", fmt.Sprintf("INV-%s%%", dateStr)).
		Find(&penjualan)

	if len(penjualan) == 0 {
		orderID = fmt.Sprintf("INV-%s-0001", dateStr)
	} else {
		orderID = fmt.Sprintf("INV-%s-%04d", dateStr, len(penjualan)+1)
	}

	return orderID
}

func CreatePenjualan(c *fiber.Ctx) error {
	var penjualan entity.OnePenjualan
	var detailPenjualan []entity.DetailPenjualan
	var user entity.User

	userID := c.Locals("user_id")
	database.DB.Preload("Toko").First(&user, "id = ?", userID)

	if err := c.BodyParser(&penjualan); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if len(penjualan.DetailPenjualan) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Detail penjualan tidak boleh kosong",
			"data":    nil,
		})
	}

	if penjualan.TotalPembayaran == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Total pembayaran tidak boleh kosong",
			"data":    nil,
		})
	}

	if penjualan.IipePembayaran == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Harap pilih tipe pembayaran",
			"data":    nil,
		})
	}

	if penjualan.IipePembayaran != "CASH" && penjualan.IipePembayaran != "QRIS" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Harap pilih tipe pembayaran CASH atau QRIS",
			"data":    nil,
		})
	}

	var total int

	for _, detail := range penjualan.DetailPenjualan {
		var produk entity.Produk
		var kategori entity.Kategori
		if err := database.DB.First(&produk, "id = ?", detail.ProdukID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Produk dengan id " + detail.ProdukID.String() + " tidak ditemukan",
				"data":    nil,
			})
		}

		if err := database.DB.First(&kategori, "id = ?", produk.KategoriID).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Kategori dengan id " + produk.KategoriID.String() + " tidak ditemukan",
				"data":    nil,
			})
		}

		if produk.Stok < detail.Quantity {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Stok produk " + produk.Nama + " tidak mencukupi",
				"data":    nil,
			})
		}

		detail.ID = uuid.New()
		detail.ProdukName = produk.Nama
		detail.Harga = produk.Harga
		detail.Subtotal = produk.Harga * detail.Quantity
		detail.Produk = produk

		total += detail.Subtotal
		detailPenjualan = append(detailPenjualan, detail)

		// update produk
		produk.Stok -= detail.Quantity
		produk.Terjual += detail.Quantity
		database.DB.Save(&produk)

		// update kategori
		kategori.Terjual += detail.Quantity
		database.DB.Save(&kategori)
	}

	penjualan.ID = uuid.New()
	penjualan.OrderID = generateOrderID()
	penjualan.TokoID = user.TokoID
	penjualan.Toko = user.Toko

	if penjualan.IipePembayaran == "QRIS" {
		penjualan.QrisID = generateQrisID()
		penjualan.Status = "UNPAID"
	} else {
		penjualan.Status = "PAID"
	}

	penjualan.UserID = user.ID
	penjualan.DetailPenjualan = detailPenjualan
	penjualan.Total = total
	penjualan.Kembalian = penjualan.TotalPembayaran - penjualan.Total
	penjualan.JumlahItem = len(detailPenjualan)

	if penjualan.TotalPembayaran < penjualan.Total {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Pembayaran tidak mencukupi",
			"data":    nil,
		})
	}

	if err := database.DB.
		Model(&entity.Penjualan{}).
		Preload("User").
		Preload("DetailPenjualan.Produk.Kategori").
		Create(&penjualan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil membuat penjualan",
		"data":    penjualan,
	})
}

func GetPenjualan(c *fiber.Ctx) error {
	tokoID := c.Locals("toko_id")

	start_date := c.Query("start_date") // YYYY-MM-DD
	end_date := c.Query("end_date")     // YYYY-MM-DD

	// pagination query
	page := utils.ParseToNumber(c.Query("page", "1"))
	limit := utils.ParseToNumber(c.Query("limit", "10"))

	// payment method query
	payment_method := c.Query("payment_method")

	var totalData int64

	var penjualan []entity.Penjualan

	analytic := entity.Analytic{
		TotalPenjualan: 0,
		TotalTransaksi: 0,
	}

	tx := database.DB.
		Model(&entity.Penjualan{}).
		Preload("User").
		Preload("DetailPenjualan.Produk").
		Where("toko_id = ?", tokoID)

	if start_date != "" && end_date != "" {
		tx = tx.Where("created_at BETWEEN ? AND ?", start_date, end_date)
	}

	if payment_method != "" {
		tx = tx.Where("tipe_pembayaran = ?", payment_method)
	}

	tx.Count(&totalData)

	// pagination
	tx = tx.Offset((page - 1) * limit).Limit(limit)

	if err := tx.Find(&penjualan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mendapatkan penjualan",
			"data":    nil,
		})
	}

	var paidPenjualan []entity.Penjualan

	if err := database.DB.Where("toko_id = ? AND status = ?", tokoID, "PAID").Find(&paidPenjualan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mendapatkan penjualan",
			"data":    nil,
		})
	}

	for _, p := range paidPenjualan {
		analytic.TotalPenjualan += p.Total
		analytic.TotalTransaksi += 1
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mendapatkan penjualan",
		"data":    penjualan,
		"summary": analytic,
		"meta": fiber.Map{
			"page":      page,
			"limit":     limit,
			"totalData": totalData,
			"totalPage": totalPage,
		},
	})
}

func GetOnePenjualan(c *fiber.Ctx) error {
	penjualanID := c.Params("id")
	var penjualan entity.OnePenjualan

	if err := database.DB.
		Model(&entity.Penjualan{}).
		Preload("User").
		Preload("Toko").
		Preload("DetailPenjualan.Produk.Kategori").
		Where("id = ?", penjualanID).
		First(&penjualan).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Penjualan dengan id " + penjualanID + " tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil mendapatkan penjualan",
		"data":    penjualan,
	})
}

func GetTopData(c *fiber.Ctx) error {
	var produk []entity.Produk
	var kategori []entity.Kategori

	var typeData = c.Query("type")

	// pagination query
	page := utils.ParseToNumber(c.Query("page", "1"))
	limit := utils.ParseToNumber(c.Query("limit", "10"))

	start_date := c.Query("start_date") // YYYY-MM-DD
	end_date := c.Query("end_date")     // YYYY-MM-DD

	if typeData != "product" && typeData != "category" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Harap pilih tipe data product atau category",
			"data":    nil,
		})
	}

	if start_date == "" || end_date == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Harap masukkan tanggal awal dan akhir",
			"data":    nil,
		})
	}

	var totalData int64

	if typeData == "product" {
		tx := database.DB.
			Model(&entity.Produk{}).
			Preload("Kategori").
			Order("terjual DESC").
			Where("terjual > ?", 0).
			Where("toko_id = ?", c.Locals("toko_id"))

		tx = tx.Where("created_at BETWEEN ? AND ?", parseDate(start_date), parseDate(end_date))

		tx.Count(&totalData)

		tx = tx.Offset((page - 1) * limit).Limit(limit)

		if err := tx.Find(&produk).Error; err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Gagal mendapatkan data produk",
				"data":    nil,
			})
		}

		totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Berhasil mendapatkan data produk",
			"data":    produk,
			"meta": fiber.Map{
				"page":      page,
				"limit":     limit,
				"totalData": totalData,
				"totalPage": totalPage,
			},
		})
	} else {
		tx := database.DB.
			Model(&entity.Kategori{}).
			Order("terjual DESC").
			Where("terjual > ?", 0).
			Where("toko_id = ?", c.Locals("toko_id"))

		tx = tx.Where("created_at BETWEEN ? AND ?", parseDate(start_date), parseDate(end_date))

		tx.Count(&totalData)

		tx = tx.Offset((page - 1) * limit).Limit(limit)

		tx.Find(&kategori)

		totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Berhasil mendapatkan data kategori",
			"data":    kategori,
			"meta": fiber.Map{
				"page":      page,
				"limit":     limit,
				"totalData": totalData,
				"totalPage": totalPage,
			},
		})
	}
}
