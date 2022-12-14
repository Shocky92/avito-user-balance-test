package main

import (
	"avito-user-balance-test/database"
	"avito-user-balance-test/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	// Application routes
	app.Get("/user/balance/:id", handlers.UserBalance)

	app.Put("/user/balance/:id/increase", handlers.IncreaseUserBalance)

	app.Post("/user/:id/order/reserve", handlers.OrderReserve)
	app.Post("/user/:id/order/proceed", handlers.OrderProceed)

	app.Listen(":3000")
}
