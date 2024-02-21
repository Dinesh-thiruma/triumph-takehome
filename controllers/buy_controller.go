package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// HandleBuy handles requests to the /buy endpoint.
func HandleBuy(c *fiber.Ctx) error {
	// Parse query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	result := services.ExecuteBuyOrder(amount, symbol)

	// Return successful response
	return c.JSON(result)
}
