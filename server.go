package main

import (
	"github.com/Dinesh-thiruma/triumph-takehome/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/buy", controllers.HandleBuy)

	app.Use(ValidateSellParams)
	app.Get("/sell", controllers.HandleSell)

	app.Listen(":4000")
}
