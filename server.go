package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/Dinesh-thiruma/triumph-takehome"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Triumph Take Home!\nBuy Endpoint: \\buy\nSell Endpoint: \\sell")
    })

	app.Get("/buy", controllers.HandleBuy)
    // app.Get("/sell", controllers.HandleSell)

    app.Listen(":3000")
}