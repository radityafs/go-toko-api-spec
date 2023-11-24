package handler

import (
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetToko(c *fiber.Ctx) error {
	var toko []entity.Toko

	err := database.DB.
		Preload("User").
		Preload("Produk").
		Find(&toko).
		Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Data berhasil ditemukan",
		"data":    toko,
	})
}

func GetOneToko(c *fiber.Ctx) error {
	id := c.Params("id")

	var toko entity.OneToko

	if err := database.DB.
		Table("toko").
		Preload("User").
		Preload("Produk.Kategori").
		Preload("Kategori").
		First(&toko, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Toko tidak ditemukan",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Data berhasil ditemukan",
		"data":    toko,
	})
}

func CreateToko(c *fiber.Ctx) error {
	request := new(request.CreateToko)
	validate := utils.NewValidator()

	if err := c.BodyParser(request); err != nil {
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

	toko := entity.Toko{
		ID:      uuid.New(),
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		Telepon: request.Telepon,
	}

	if err := database.DB.Create(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  201,
		"message": "Toko berhasil dibuat",
		"data":    toko,
	})
}

func UpdateToko(c *fiber.Ctx) error {
	var request request.UpdateToko
	validate := utils.NewValidator()

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Gagal memperbarui toko",
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

	id := c.Params("id")

	var toko entity.Toko

	if err := database.DB.First(&toko, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Toko tidak ditemukan",
			"data":    nil,
		})
	}

	if err := database.DB.Model(&toko).Updates(request).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Gagal memperbarui toko",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil memperbarui toko",
		"data":    toko,
	})
}

func DeleteToko(c *fiber.Ctx) error {
	id := c.Params("id")

	var toko entity.Toko

	if err := database.DB.First(&toko, "id = ?", id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Toko tidak ditemukan",
			"data":    nil,
		})
	}

	if err := database.DB.Delete(&toko).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Gagal menghapus toko",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Berhasil menghapus toko",
		"data":    nil,
	})
}
