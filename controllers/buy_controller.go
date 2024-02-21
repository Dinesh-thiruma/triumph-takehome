package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/your-username/go-crypto-server/services"
)

// HandleBuy handles requests to the /buy endpoint.
func HandleBuy(c *fiber.Ctx) error {
    // Parse query parameters
    amount := c.Query("amount")
    symbol := c.Query("symbol")

	return c.SendString("Received amount: " + amount + ", symbol: " + symbol)
}