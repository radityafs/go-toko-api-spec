package handler

import (
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/model/response"
	"go-toko/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetUser(ctx *fiber.Ctx) error {
	var users []entity.User

	err := database.DB.
		Preload("Toko").
		Find(&users).
		Error

	if err != nil {
		return ctx.Status(500).JSON(response.User{
			Status:  500,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(200).JSON(response.User{
		Status:  200,
		Message: "Data berhasil ditemukan",
		Data:    users,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	request := new(request.CreateUser)
	validate := utils.NewValidator()

	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(400).JSON(response.User{
			Status:  400,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := validate.Struct(request); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": utils.ValidatorErrors(err),
			"data":    nil,
		})
	}

	user := entity.User{
		ID:     uuid.New(),
		Email:  request.Email,
		Nama:   request.Nama,
		Role:   request.Role,
		TokoID: request.TokoID,
	}

	hased, errHash := utils.HashPassword(request.Password)

	if errHash != nil {
		return ctx.Status(500).JSON(response.User{
			Status:  500,
			Message: errHash.Error(),
			Data:    nil,
		})
	}

	user.Password = hased

	err := database.DB.Create(&user).Error

	if err != nil {
		return ctx.Status(500).JSON(response.User{
			Status:  500,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.Status(201).JSON(response.DetailUser{
		Status:  201,
		Message: "User berhasil dibuat",
		Data:    user,
	})
}
