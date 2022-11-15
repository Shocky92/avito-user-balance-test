package handlers

import (
	"avito-user-balance-test/database"
	"avito-user-balance-test/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserBalance(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var userBalance models.UserBalance

	result := db.Find(&userBalance, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&userBalance)
}

func IncreaseUserBalance(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var userBalance models.UserBalance

	err := db.Find(&userBalance, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUserBalance := new(models.UserBalance)

		if err := c.BodyParser(newUserBalance); err != nil {
			return c.SendStatus(500)
		}

		db.Create(newUserBalance)

		return c.Status(201).JSON(&newUserBalance)
	}

	updateUserBalance := new(models.UserBalance)

	if err := c.BodyParser(updateUserBalance); err != nil {
		return c.SendStatus(500)
	}

	userBalance.Balance += updateUserBalance.Balance

	db.Save(&userBalance)

	return c.Status(200).JSON(&userBalance)
}

func OrderReserve(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var userBalance models.UserBalance

	err := db.Find(&userBalance, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}

	userOrder := new(models.UserOrder)

	if err := c.BodyParser(userOrder); err != nil {
		return c.SendStatus(500)
	}

	// check if user have enought money for order
	if userBalance.Balance < userOrder.Cost {
		return c.SendStatus(400)
	}

	userBalance.Balance -= userOrder.Cost
	db.Create(userOrder)
	userOrder.IsReserved = true
	db.Save(userBalance)

	return c.Status(201).JSON(&userOrder)
}

func OrderProceed(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	var userOrder models.UserOrder

	userOrderProceed := new(models.UserOrder)
	if err := c.BodyParser(userOrderProceed); err != nil {
		return c.SendStatus(500)
	}

	err := db.Where(
		"user_id = ? AND service_id = ? AND order_id = ? AND cost = ?",
		id,
		userOrderProceed.ServiceId,
		userOrderProceed.OrderId,
		userOrderProceed.Cost,
	).Find(&userOrder).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}

	userOrder.IsReserved = false
	db.Save(&userOrder)

	return c.Status(200).JSON(&userOrder)
}
