package controllers

import (
	"strconv"

	"github.com/Dinesh-thiruma/triumph-takehome/services"
	"github.com/gofiber/fiber/v2"
)

// ValidateSellParams validates the amount and symbol query parameters for the /sell endpoint.
func ValidateSellParams(c *fiber.Ctx) error {
	// Parse query parameters
	amount := c.Query("amount")

	// Validate amount parameter
	amountInt, err := strconv.Atoi(amount)
	if err != nil || amountInt <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid amount parameter. Amount must be a positive integer.",
		})
	}

	// You can add further validation for the symbol parameter if needed
	// For now, we assume any string is valid for symbol

	return c.Next()
}

// HandleBuy handles requests to the /buy endpoint.
func HandleSell(c *fiber.Ctx) error {
	// Parse query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	resp := services.GetAverageSell(amount, symbol)

	return c.JSON(resp)
}
