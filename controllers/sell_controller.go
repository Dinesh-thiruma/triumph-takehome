package controllers

import (
	"github.com/Dinesh-thiruma/triumph-takehome/services"
	"github.com/gofiber/fiber/v2"
)

// HandleBuy handles requests to the /buy endpoint.
func HandleSell(c *fiber.Ctx) error {
	// Parse query parameters
	amount := c.Query("amount")
	symbol := c.Query("symbol")

	resp := services.GetAverageSell(amount, symbol)

	return c.JSON(resp)
}
