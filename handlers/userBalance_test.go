package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserBalance(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "get HTTP status 200",
			route:        "/user/balance/1",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when user balance is not exists",
			route:        "/user/balance/not-exists",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	app.Get("/user/balance/1", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"UserId":  1,
			"balance": 200,
		})
	})

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, test.route, nil)
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestIncreaseUserBalance(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		requestBody  string
	}{
		{
			description:  "get HTTP status 200",
			route:        "/user/balance/1/increase",
			expectedCode: 200,
			requestBody:  "UserId=1&Balance=200",
		},
		{
			description:  "get HTTP status 201, when new user increase balance",
			route:        "/user/balance/2/increase",
			expectedCode: 201,
			requestBody:  "UserId=2&Balance=200",
		},
	}

	app := fiber.New()

	app.Put("/user/balance/1/increase", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"UserId":  1,
			"balance": 200,
		})
	})

	app.Put("/user/balance/2/increase", func(c *fiber.Ctx) error {
		return c.Status(201).JSON(&fiber.Map{
			"UserId":  2,
			"balance": 200,
		})
	})

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPut, test.route, strings.NewReader(test.requestBody))
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestOrderReserve(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		requestBody  string
	}{
		{
			description:  "get HTTP status 200",
			route:        "/user/1/order/reserve",
			expectedCode: 200,
			requestBody:  "UserId=1&ServiceId=1&OrderId=1&Cost=200&IsReserved=1",
		},
	}

	app := fiber.New()

	app.Post("/user/1/order/reserve", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"UserId":     1,
			"ServiceId":  100,
			"OrderId":    900,
			"Cost":       200,
			"IsReserved": true,
		})
	})

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, test.route, strings.NewReader(test.requestBody))
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestOrderProceed(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
		requestBody  string
	}{
		{
			description:  "get HTTP status 200",
			route:        "/user/1/order/proceed",
			expectedCode: 200,
			requestBody:  "UserId=1&ServiceId=1&OrderId=1&Cost=200&IsReserved=0",
		},
	}

	app := fiber.New()

	app.Post("/user/1/order/proceed", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"UserId":     1,
			"ServiceId":  100,
			"OrderId":    900,
			"Cost":       200,
			"IsReserved": false,
		})
	})

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodPost, test.route, strings.NewReader(test.requestBody))
		resp, _ := app.Test(req, 1)
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
