package handler

import (
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetKategori(c *fiber.Ctx) error {
	var kategori []entity.Kategori
	tokoID := c.Locals("toko_id")

	var totalData int64

	tx := database.DB.
		Where("toko_id = ?", tokoID).
		Find(&kategori)

	if tx.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": tx.Error.Error(),
			"data":    nil,
		})
	}

	tx.Count(&totalData)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Data berhasil ditemukan",
		"data":    kategori,
	})
}

func CreateKategori(c *fiber.Ctx) error {
	var request request.CreateKategori
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

	kategori := entity.Kategori{
		ID:     uuid.New(),
		Nama:   request.Nama,
		TokoID: request.TokoID,
	}

	if err := database.DB.Create(&kategori).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Gagal membuat kategori",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil membuat kategori",
		"data":    kategori,
	})
}
