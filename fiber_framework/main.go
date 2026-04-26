package main

import (
	"fiber-framework/config"
	"fiber-framework/handlers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.ConnectDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Post("/cars", handlers.CreateCars)
	app.Put("/cars/:id", handlers.UpdateCar)
	app.Delete("/cars/:id", handlers.DeleteCar)
	app.Get("/cars", handlers.GetAllCars)
	app.Get("/cars/:id", handlers.GetById)

	fmt.Println("server started")
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println("server stopped abruptly..")
		return
	}
}
