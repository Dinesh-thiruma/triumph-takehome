package controllers

import (
	"github.com/Dinesh-thiruma/triumph-takehome/services"
	"github.com/gofiber/fiber/v2"
)

// HandleBuy handles requests to the /buy endpoint.
func HandleBuy(c *fiber.Ctx) error {
	// Parse query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	resp := services.GetAverage(amount, symbol)

	return c.JSON(resp)
}
