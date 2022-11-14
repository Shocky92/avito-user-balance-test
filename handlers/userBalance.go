package handlers

import (
	"avito-user-balance-test/database"
	"avito-user-balance-test/models"

	"github.com/gofiber/fiber/v2"
)

func UserBalance(c *fiber.Ctx) error {
	id := c.Params("id")
	var userBalance models.UserBalance

	result := database.DB.Db.Find(&userBalance, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&userBalance)
}
