package handler

import (
	"fmt"
	"go-toko/database"
	"go-toko/model/entity"
	"go-toko/model/request"
	"go-toko/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Login(ctx *fiber.Ctx) error {
	loginRequest := new(request.Login)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": utils.ValidatorErrors(err),
			"data":    nil,
		})
	}

	var user entity.User
	if err := database.DB.
		Preload("Toko").
		First(&user, "email = ?", loginRequest.Email).Error; err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User dengan email yang diberikan tidak ditemukan",
			"data":    nil,
		})
	}

	if passed := utils.CheckPasswordHash(loginRequest.Password, user.Password); !passed {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password yang diberikan salah",
			"data":    nil,
		})
	}

	// TODO: Generate JWT Token
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["toko_id"] = user.TokoID
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 1 day

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = user.ID
	refreshClaims["email"] = user.Email
	refreshClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix() // 7 days

	token, err := utils.GenerateToken(&claims)
	refreshToken, errRefresh := utils.GenerateToken(&refreshClaims)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
			"data":    nil,
		})
	}

	if errRefresh != nil {
		fmt.Println(errRefresh)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success":      true,
		"message":      "Successfuly logged in",
		"data":         user,
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func RefreshToken(ctx *fiber.Ctx) error {
	var token request.RefreshToken
	var user entity.User

	if err := ctx.BodyParser(&token); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	claims, err := utils.VerifyAndParseToken(token.RefreshToken)

	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Token tidak valid",
			"data":    nil,
		})
	}

	now := time.Now().Unix()
	if claims["exp"] == nil || int64(claims["exp"].(float64)) < now {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Token expired",
			"data":    nil,
		})
	}

	if claims["user_id"] == nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Token tidak valid",
			"data":    nil,
		})
	}

	if err := database.DB.First(&user, "id = ?", claims["user_id"]).Error; err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "User tidak ditemukan",
			"data":    nil,
		})
	}

	newClaims := jwt.MapClaims{}
	newClaims["user_id"] = user.ID
	newClaims["toko_id"] = user.TokoID
	newClaims["email"] = user.Email
	newClaims["role"] = user.Role
	newClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 1 day

	newToken, err := utils.GenerateToken(&newClaims)

	if err != nil {
		fmt.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
			"data":    nil,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Successfuly refreshed token",
		"token":   newToken,
	})
}
