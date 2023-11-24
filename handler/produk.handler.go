package handler

import (
	"context"
	"fmt"
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/utils"
	"math"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func parseSortBy(sort string) string {
	switch sort {
	case "name":
		return "nama"
	case "price":
		return "harga"
	default:
		return "created_at"
	}
}

var CLOUDINARY_URL = "cloudinary://457255529371718:mMmvp9ch54Txch8Iews5veWXdCs@dh9qcux3p"

func CreateProduk(c *fiber.Ctx) error {
	var request request.CreateProduk
	validate := utils.NewValidator()

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err := validate.Struct(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": utils.ValidatorErrors(err),
			"data":    nil,
		})
	}

	fileHeader, err := c.FormFile("photo")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
			"data":    nil,
		})
	}

	file, _ := fileHeader.Open()

	ctx := context.Background()

	cldService, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
	resp, errUpload := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	if errUpload != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Gagal upload image",
			"data":    nil,
		})
	}

	produk := entity.Produk{
		ID:         uuid.New(),
		Nama:       request.Nama,
		Image:      resp.SecureURL,
		TokoID:     request.TokoID,
		Stok:       request.Stok,
		KategoriID: request.KategoriID,
	}

	harga, err := strconv.Atoi(c.FormValue("price"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Harga harus berupa angka",
			"data":    nil,
		})
	}

	produk.Harga = harga

	if err := database.DB.
		Preload("Kategori").
		Create(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Gagal membuat produk",
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  201,
		"message": "Berhasil membuat produk",
		"data":    produk,
	})
}

func GetProduk(c *fiber.Ctx) error {
	tokoId := c.Locals("toko_id")

	var totalData int64

	kategoriId := c.Query("id_category")
	nama := c.Query("query")

	// pagination query
	page := utils.ParseToNumber(c.Query("page", "1"))
	limit := utils.ParseToNumber(c.Query("limit", "10"))

	// sort query
	sort := c.Query("sort")
	order := c.Query("order")

	var produk []entity.Produk

	tx := database.DB.
		Model(&entity.Produk{}).
		Preload("Toko").
		Preload("Kategori").
		Where("toko_id = ?", tokoId)

	if kategoriId != "" {
		tx = tx.Where("kategori_id = ?", kategoriId)
	}

	if nama != "" {
		tx = tx.Where("nama LIKE ?", "%"+nama+"%")
	}

	if sort != "" && order != "" {
		tx = tx.Order(parseSortBy(sort) + " " + strings.ToUpper(order))
	} else {
		tx = tx.Order("created_at DESC")
	}

	tx.Count(&totalData)

	// pagination
	tx = tx.Offset((page - 1) * limit).Limit(limit)

	if err := tx.Find(&produk).Error; err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal mendapatkan produk",
			"data":    nil,
		})
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil mendapatkan produk",
		"data":    produk,
		"meta": fiber.Map{
			"page":      page,
			"limit":     limit,
			"totalData": totalData,
			"totalPage": totalPage,
		},
	})
}

func UpdateProduk(c *fiber.Ctx) error {
	var request request.UpdateProduk
	validate := utils.NewValidator()

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Gagal memperbarui produk",
			"data":    nil,
		})
	}

	if err := validate.Struct(request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": utils.ValidatorErrors(err),
			"data":    nil,
		})
	}

	id := c.Params("id")
	tokoId := c.Locals("toko_id")

	var produk entity.Produk

	if err := database.DB.
		Preload("Kategori").
		Where("id = ? AND toko_id = ?", id, tokoId).
		First(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Produk tidak ditemukan",
			"data":    nil,
		})
	}

	haveImage, _ := c.FormFile("photo")

	if haveImage != nil {
		fileHeader, err := c.FormFile("photo")

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
				"data":    nil,
			})
		}

		file, _ := fileHeader.Open()

		ctx := context.Background()

		cldService, _ := cloudinary.NewFromURL(CLOUDINARY_URL)
		resp, errUpload := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

		if errUpload != nil {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Gagal upload image",
				"data":    nil,
			})
		}

		produk.Image = resp.SecureURL

		if err := database.DB.Save(&produk).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"success": false,
				"message": "Gagal memperbarui produk",
				"data":    nil,
			})
		}
	}

	if err := database.DB.Model(&produk).
		Preload("Kategori").
		Updates(request).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Gagal memperbarui produk",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Berhasil memperbarui produk",
		"data":    produk,
	})
}

func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")

	var produk entity.Produk

	if err := database.DB.
		Where("id = ?", id).
		First(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Produk tidak ditemukan",
			"data":    nil,
		})
	}

	if err := database.DB.Delete(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Gagal menghapus produk",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil menghapus produk",
		"data":    produk,
	})
}
