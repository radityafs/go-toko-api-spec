package middleware

import (
	"go-toko/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	bearerSchema := c.Get("Authorization")

	if bearerSchema == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, token tidak ditemukan",
			"data":    nil,
		})
	}

	splitToken := strings.Split(bearerSchema, " ")
	if len(splitToken) != 2 {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, token tidak valid",
			"data":    nil,
		})
	}

	token := splitToken[1]

	claims, err := utils.VerifyAndParseToken(token)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, token tidak valid",
			"data":    nil,
		})
	}

	now := time.Now().Unix()
	if claims["exp"] == nil || int64(claims["exp"].(float64)) < now {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized, token expired",
			"data":    nil,
		})
	}

	c.Locals("user_id", claims["user_id"])
	c.Locals("toko_id", claims["toko_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden, anda tidak memiliki akses",
			"data":    nil,
		})
	}

	return c.Next()
}

func IsOwner(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "owner" {
		return c.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden, anda tidak memiliki akses",
			"data":    nil,
		})
	}

	return c.Next()
}

func PermissionCreate(c *fiber.Ctx) error {

	return c.Next()
}
